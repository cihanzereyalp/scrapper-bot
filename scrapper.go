package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func Scrape(url string, c chan string, pattern string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	r, _ := regexp.Compile(pattern)
	res := r.FindStringSubmatch(sb)
	if res != nil {
		//c <- res[1] + url
		c <- "Players: " + res[1]
		return
	}
	c <- "Not found" + url

}
