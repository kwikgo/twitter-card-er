package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
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
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.Black)
			default:
				// Use zero value.
			}
		}
	}

	// Party sticker
	partyPosition := image.Point{0, 0}
	// Niemals Fehler zu erwarten bei der Partei! Embed, wissen's schon!
	dieParteiLogo, _ := jpeg.Decode(bytes.NewReader(diePartei))
	partyRectangle := image.Rectangle{partyPosition, partyPosition.Add(dieParteiLogo.Bounds().Size())}
	draw.Draw(img, partyRectangle, dieParteiLogo, image.Point{0, 0}, draw.Src)

	buf := &bytes.Buffer{}
	png.Encode(buf, img)
	return buf
}

func loadBackDrop(path string) {
	// if startsWith http:// load remote if valid png or jpeg
	//return *Image
}
