package main

import "net/http"

type Card struct {
	Title           string
	Description     string
	Site            string
	BaseURL         string
	TextColor       string
	BackgroundColor string
	Backdrop        string
	QueryString     string
}

func urlParamsToCard(r *http.Request) Card {
	return Card{
		Title:           thisOrThat(r.URL.Query().Get("title"), defaults.Title),
		Description:     thisOrThat(r.URL.Query().Get("description"), defaults.Description),
		Site:            thisOrThat(r.URL.Query().Get("site"), defaults.Site),
		BaseURL:         thisOrThat(r.URL.Query().Get("baseurl"), defaults.BaseURL),
		TextColor:       thisOrThat(r.URL.Query().Get("textcolor"), defaults.TextColor),
		BackgroundColor: thisOrThat(r.URL.Query().Get("backgroundcolor"), defaults.Backdrop),
		Backdrop:        thisOrThat(r.URL.Query().Get("backdrop"), defaults.Backdrop),
		QueryString:     r.URL.RawQuery,
	}
}

func thisOrThat(this, that string) string {
	if this == "" {
		return that
	}
	return this
}

func (card *Card) hash() string {
	return "base64 for url params - "
}
