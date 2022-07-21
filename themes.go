package main

import "log"

func (card *Card) applyTheme() {
	themeCard, themeExists := Themes[card.Theme]
	if card.Theme == "" || !themeExists {
		card.BackgroundColor = themeCard.BackgroundColor
		return
	}
	log.Printf("applying custom theme with bgc: %s", themeCard.BackgroundColor)
}

func loadThemes() {
	Themes = make(map[string]Card)

	Themes["drucksache"] = Card{
		Title:           "Drucksache",
		BackgroundColor: "todo",
	}

	Themes["duden"] = Card{
		Title: "FAKE-Duden",
	}
}
