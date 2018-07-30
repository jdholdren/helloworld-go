package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
)

const (
	urlName = "SLACK_URL"
)

func handler(w http.ResponseWriter, r *http.Request) {
	webHookURL := os.Getenv(urlName)
	msg := "Hello, world!"

	payload := slack.Payload{
		Text:      msg,
		Username:  "robot",
		Channel:   "#general",
		IconEmoji: ":monkey_face:",
	}

	// Send
	err := slack.Send(webHookURL, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
