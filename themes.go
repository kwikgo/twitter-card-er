package main

func (card *Card) applyTheme() {
	themeCard, themeExists := Themes[card.Theme]
	if card.Theme == "" || !themeExists {
		card.BackgroundColor = themeCard.BackgroundColor
		return
	}
}

func loadThemes() {
	Themes = make(map[string]Card)

	Themes["drucksache"] = Card{
		Title:           "Drucksache",
		BackgroundColor: "drucksacherot",
	}

	Themes["duden"] = Card{
		Title:           "FAKE-Duden",
		BackgroundColor: "dudengelb",
	}
}
