package slackbot

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

const myToken = "hF5F1iTzGUUFWDI8gnS0JPIy"

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

func init() {
	hi := http.HandlerFunc(Hi)

	http.Handle("/hi", AuthHandler(hi))
	http.HandleFunc("/bad_advice", BadAdviceHanlder)
}

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			form := &SlashForm{}
			decoder := schema.NewDecoder()
			r.ParseForm()
			token := r.FormValue("token")

			err := decoder.Decode(form, r.PostForm)
			if err != nil {
				return
			}

			if token != myToken {
				w.WriteHeader(404)
				w.Write([]byte("Invalid Token"))
			}
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("bad request"))
		}
	})
}

func Hi(w http.ResponseWriter, r *http.Request) {
	var response *SlashResponse
	response = &SlashResponse{ResponseType: "in_channel", Text: "This is a greeting"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func BadAdviceHanlder(w http.ResponseWriter, r *http.Request) {
	var response *SlashResponse
	response = &SlashResponse{ResponseType: "in_channel", Text: BadAdvice()}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
