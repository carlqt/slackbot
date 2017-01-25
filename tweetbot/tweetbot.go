package tweetbot

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

var consumerKey, secret string
var request *http.Request

const baseURL = "https://api.twitter.com/"

type twitterResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type UserStatus struct {
	Text string `json:"text"`
	// ID   string `json:"id"`
}

func init() {
	consumerKey = "uC0SUjvcC7vdO9nmDIrfoJSFq"
	secret = "6xI88YRMQIU2yAWhu9yR5ZsqwtzKTmpxiZ4ETXZzS753JsLc1h"
}

func TweetTCL() string {
	twitterUser := map[string]string{"thecodinglove": "929531611", "lifeadvicelamp": "3407769339"}

	token := getBearerToken()
	return getTweet(twitterUser["thecodinglove"], token)
}

func getBearerToken() string {
	twitResp := twitterResponse{}
	authURL := baseURL + "oauth2/token"
	client := &http.Client{}
	// client := &http.Client{}

	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	encodedKey := encode()

	req, _ := http.NewRequest("POST", authURL, bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", "Basic "+encodedKey)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&twitResp)
	return twitResp.AccessToken
}

func getTweet(userID string, token string) string {
	tweets := []UserStatus{}

	// Call user_timeline twitterAPI and store json reponse in tweets struct
	searchURL := baseURL + "1.1/statuses/user_timeline.json"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", searchURL, nil)
	q := req.URL.Query()
	q.Add("count", "500")
	q.Add("exclude_replies", "true")
	q.Add("user_id", userID)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+token)

	log.Println("Connecting to " + searchURL)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	json.NewDecoder(resp.Body).Decode(&tweets)

	tweet := getRandomTweet(tweets)
	return FormatTweet(tweet)
}

func getRandomTweet(tweets []UserStatus) string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(tweets))
	return fmt.Sprint(tweets[n].Text)
}

func encode() (enc string) {
	str := consumerKey + ":" + secret
	enc = base64.StdEncoding.EncodeToString([]byte(str))
	return enc
}

// for debugging purposes
func debugResponse(response io.ReadCloser) {
	io.Copy(os.Stdout, response)
}

func debugRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}
