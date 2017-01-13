package main

import(
  "fmt"
  // "github.com/nlopes/slack"
  "net/http"
  "encoding/json"
)

const token = "xoxb-126064414341-lXGF4TYMqaH4aNEUJBoGPFw1"
const carlAppToken = "xoxp-125270328032-126053736308-127158113959-fa434655026cdb8a4fa0026896399a78"

type slackUserList struct {
  Ok string
}

func main() {
  jsonResp := slackUserList{}

  url := fmt.Sprintf("https://slack.com/api/users.list?token=%s&pretty", carlAppToken)

  resp, err := http.Get(url)
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()
  json.NewDecoder(resp.Body).Decode(&jsonResp)
  fmt.Println(jsonResp)
}
