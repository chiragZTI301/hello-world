package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func img() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	createAndSaveImg()
}

func createAndSaveImg() {
	width, height := 300, 200
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Set the background color to white
	backgroundColor := color.RGBA{255, 255, 255, 255}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, backgroundColor)
		}
	}

	// Set a blue rectangle in the middle
	rectangleColor := color.RGBA{0, 0, 255, 255}
	for y := height / 2; y < 3*height/4; y++ {
		for x := width / 4; x < 3*width/4; x++ {
			img.Set(x, y, rectangleColor)
		}
	}

	// Save the image to a file
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}