package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

type TestP struct {
	Token string
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func Hi(w http.ResponseWriter, r *http.Request) {
	var response *SlashResponse
	form := &SlashForm{}
	decoder := schema.NewDecoder()
	r.ParseForm()

	err := decoder.Decode(form, r.PostForm)
	if err != nil {
		panic(err)
	}
	response = &SlashResponse{ResponseType: "in_channel", Text: "This is a greeting"}

	token := r.FormValue("token")
	myToken := "hF5F1iTzGUUFWDI8gnS0JPIy"

	w.Header().Set("Content-Type", "application/json")

	if token != myToken {
		fmt.Fprint(w, "Invalid Token")
	} else {
		w.WriteHeader(200)

		// fmt.Fprint(w, "hello")
		json.NewEncoder(w).Encode(response)
	}
}
