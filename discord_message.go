package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DiscordMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func (m *DiscordMessage) Send(webhookURL string) error {
	resultBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(resultBytes))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
