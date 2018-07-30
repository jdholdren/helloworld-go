package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	urlName = "SLACK_URL"
)

func handler(w http.ResponseWriter, r *http.Request) {
	webHookURL := os.Getenv(urlName)
	if webHookURL == "" {
		writeError(errors.New("Webhook url is not defined"), w)
		return
	}

	msg := "Hello, world!"

	body := struct {
		Text string `json:"text"`
	}{
		Text: msg,
	}

	// Marshall the body to json
	bBytes, err := json.Marshal(body)
	if err != nil {
		writeError(err, w)
		return
	}

	// Make the request
	req, err := http.NewRequest("POST", webHookURL, bytes.NewBuffer(bBytes))
	if err != nil {
		writeError(err, w)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Open the client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		writeError(err, w)
		return
	}
	defer resp.Body.Close()

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writeError(err, w)
		return
	}

	// Successful request
	w.WriteHeader(200)
	w.Write(rBody)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeError(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
