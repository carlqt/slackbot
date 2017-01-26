package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/carlqt/slackbot/tweetbot"

	"github.com/gorilla/schema"
)

type SlashForm struct {
	Token       string `schema:"token"`
	TeamID      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelID   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	UserID      string `schema:"user_id"`
	Command     string
	Text        string
	UserName    string `schema:"user_name"`
	ResponseURL string `schema:"response_url"`
}

type SlashResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html") // Parse template file.

	if err != nil {
		log.Println(err)
	} else {
		t.Execute(w, nil) // merge.
	}
}

func TokenHandler(next http.Handler, validToken string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form := &SlashForm{}
		decoder := schema.NewDecoder()
		r.ParseForm()
		token := r.FormValue("token")

		err := decoder.Decode(form, r.PostForm)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Invalid Token"))
			return
		}

		if token != validToken && r.Method == "POST" {
			w.WriteHeader(404)
			w.Write([]byte("Invalid Token"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Tweet(w http.ResponseWriter, r *http.Request) {
	defaultResponse := &SlashResponse{ResponseType: "ephemeral", Text: "success"}

	// prepare to make a slack delayed response
	responseURL := r.FormValue("response_url")
	go sendDelayResponse(responseURL)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(defaultResponse)
}

func sendDelayResponse(url string) {
	response := &SlashResponse{ResponseType: "in_channel", Text: tweetbot.TweetTCL()}
	b, _ := json.Marshal(response)

	// send request to slack using given response_url
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	log.Println("Connecting to " + url)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
	} else {
		log.Println("success")
	}
}
