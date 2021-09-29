package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type LineNotification struct {
	Message string
}

func (line LineNotification) send() {
	token := "" // token発行して追加してください
	uri := "https://notify-api.line.me/api/notify"
	data := url.Values{}
	data.Set("message", line.Message)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Print("Line message was sent")
}

func main() {
	lineNotification := LineNotification{Message: "Line message sent by okubo."}
	lineNotification.send()
}
