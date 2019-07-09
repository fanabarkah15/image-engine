package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "http://i.imgur.com/m1UIjW1.jpg"
	url2 := "https://png.pngtree.com/element_pic/17/08/10/a9d96c533d2304109902cd8aed58d98b.jpg"

	response, _ := http.Get(url)

	defer response.Body.Close()
	first, _, err := image.Decode(response.Body)

	response, _ = http.Get(url2)

	defer response.Body.Close()
	second, _, err := image.Decode(response.Body)

	offset := image.Pt(300, 200)
	b := first.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.ZP, draw.Src)
	draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)

	third, err := os.Create("result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()

}
