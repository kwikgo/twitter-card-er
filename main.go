package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"
)

var (
	defaults Card
	Themes   map[string]Card
)

func main() {
	loadThemes()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/png/", pngHandler)

	defaults.Title = thisOrThat(os.Getenv("TCE_TITLE"), "The DEFAULT title for twitter-card-er. Provide it!")
	defaults.Site = thisOrThat(os.Getenv("TCE_SITE"), "@JanHacker9")
	defaults.BaseURL = thisOrThat(os.Getenv("TCE_BASEURL"), "https://hacker.ch/twitter-card-er")
	defaults.Domain = thisOrThat(os.Getenv("TCE_DOMAIN"), "hacker.ch") // note - displayed only 1x in text below img

	err := http.ListenAndServe(":9911", nil) // fixme env
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
