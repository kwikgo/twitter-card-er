package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"

	"github.com/fogleman/gg"
)

var (
	//go:embed DiePARTEI.jpg
	diePartei []byte
)

func mkImage(card Card) *bytes.Buffer {
	// if validCacheExists(card.hash()) {
	// use hashed urlparams as cachekey
	// keep last 10 or so. housekeeping... :/
	// }

	width := 1120
	height := 600

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	bgColor := getBackgroundColor(card.Theme, card.BackgroundColor)
	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, bgColor)
		}
	}

	// Party sticker
	partyPosition := image.Point{0, 0}
	// Niemals Fehler zu erwarten bei der Partei! Embed, wissen's schon!
	dieParteiLogo, _ := jpeg.Decode(bytes.NewReader(diePartei))
	partyRectangle := image.Rectangle{partyPosition, partyPosition.Add(dieParteiLogo.Bounds().Size())}
	draw.Draw(img, partyRectangle, dieParteiLogo, image.Point{0, 0}, draw.Src)

	// via https://sourcegraph.com/github.com/fogleman/gg/-/blob/examples/gofont.go
	font, err := truetype.Parse(gobold.TTF)
	if err != nil {
		log.Fatal(err)
	}
	face := truetype.NewFace(font, &truetype.Options{Size: 52})
	dc := gg.NewContextForImage(img)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(card.Title, 560, 420, 0.5, 0.5)
	dc.SetRGB(0.7, 0, 0)
	dc.DrawStringAnchored(card.Description, 560, 500, 0.5, 0.5)

	buf := &bytes.Buffer{}
	png.Encode(buf, dc.Image())
	return buf
}

func getBackgroundColor(theme, namedColor string) color.Color {
	targetColor := "black"
	tc := color.RGBA{0, 0, 0, 0}

	if theme != "" {
		targetColor = Themes[theme].BackgroundColor
	}
	if namedColor != "" {
		targetColor = namedColor
	}
	switch targetColor {
	case "cyan":
		tc = color.RGBA{100, 200, 200, 0xff} // cyan, default - nice!
	case "drucksacherot":
		tc = color.RGBA{164, 12, 33, 0xff}
	case "dudengelb":
		tc = color.RGBA{254, 194, 10, 0xff}
	}
	return tc
}

func loadBackDrop(path string) {
	// if startsWith http:// load remote if valid png or jpeg
	//return *Image
}
