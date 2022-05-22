package main

import "testing"

func TestCreateMessage(t *testing.T) {
	got := CreateMessage(Scrapper{
		[]string{`https://webscraper.io/test-sites/e-commerce/allinone-popup-links`},
		`yourWebHookURL`,
		`<p>.*?\n.*?Welcome to (.*?) site\.`,
	})
	want := []string{`Players: WebScraper e-commerce`}

	if got[0] != want[0] {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
