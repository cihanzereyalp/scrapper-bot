package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Scrapper struct {
	URLList    []string `json:"urlList"`
	WebHookURL string   `json:"webHookURL"`
	Pattern    string   `json:"pattern"`
}

func main() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("opening config file", err)
		return
	}
	defer configFile.Close()
	jsonData, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}
	scrapper := Scrapper{}
	if err := json.Unmarshal(jsonData, &scrapper); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return
	}
	resultList := CreateMessage(scrapper)
	m := DiscordMessage{
		"bot",
		strings.Join(resultList, "\n"),
	}
	m.Send(scrapper.WebHookURL)
}

func CreateMessage(scrapper Scrapper) []string {
	c := make(chan string, len(scrapper.URLList))
	for _, url := range scrapper.URLList {
		go Scrape(url, c, scrapper.Pattern)
	}
	var resultList []string
	for _, _ = range scrapper.URLList {
		resultList = append(resultList, <-c)
	}
	return resultList
}
