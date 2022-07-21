package main

import (
	"bytes"
	_ "embed"
	"log"
	"net/http"
	"text/template"
)

var (
	//go:embed card_template.html
	indexHTML string
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	card := urlParamsToCard(r)
	w.Write(renderIndexTemplate(card))
}

func pngHandler(w http.ResponseWriter, r *http.Request) {
	card := urlParamsToCard(r)
	card.applyTheme()
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
