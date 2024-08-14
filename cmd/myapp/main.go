package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("./images/input.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(i)
	bounds := img.Bounds()

	grayimg := image.NewGray(bounds)
	newimg := image.NewRGBA(bounds)

	outcolor, err := os.Create("./images/outcolor.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outcolor.Close()

	outgray, err := os.Create("./images/outgray.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outgray.Close()

	err = png.Encode(outgray, grayimg)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(outcolor, newimg)
	if err != nil {
		fmt.Println(err)
		return
	}

}
