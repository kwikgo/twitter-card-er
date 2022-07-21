package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	//go:embed index.html
	indexHTML string
	defaults  Card
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/png/", pngHandler)

	defaults.Title = thisOrThat(os.Getenv("TCE_TITLE"), "The DEFAULT title for twitter-card-er. Provide it!")
	defaults.Site = thisOrThat(os.Getenv("TCE_SITE"), "@JanHacker9")
	defaults.BaseURL = thisOrThat(os.Getenv("TCE_BASEURL"), "https://hacker.ch/twitter-card-er")

	err := http.ListenAndServe(":9911", nil) // fixme env
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	card := urlParamsToCard(r)
	w.Write(renderIndexTemplate(card))
}

func pngHandler(w http.ResponseWriter, r *http.Request) {
	card := urlParamsToCard(r)
	png := mkImage(card)
	w.Header().Add("content-type", "image/png")
	w.Write(png.Bytes())
}

func renderIndexTemplate(card Card) []byte {
	buf := &bytes.Buffer{}
	tpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		log.Fatalf("Template parsing error: %v\n", err)
	}
	err = tpl.Execute(buf, card)
	if err != nil {
		log.Printf("Template execution error: %v\n", err)
	}
	return buf.Bytes()
}
