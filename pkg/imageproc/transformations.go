package utils

import (
	"fmt"
	"image"
	"image/color"
)

var ARRAY = []string{" ", ".", "-", "*", "%"}
var LETTERS = []string{" ", ".", "-", "'", ":", "_", ",", "^", "=", ";", ">", "<", "+", "!", "r", "c", "*", "/", "z", "?", "R", "D", "#", "$", "B", "g", "0", "M", "N", "W", "Q", "%", "&", "@"}

func InvertColors(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newimg := image.NewRGBA(bounds)

	w, h := bounds.Max.X, bounds.Max.Y

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := img.At(x, y)
			r, g, b, a := pix.RGBA()

			newpix := color.RGBA{
				R: uint8(255 - r>>8),
				G: uint8(255 - g>>8),
				B: uint8(255 - b>>8),
				A: uint8(a >> 8),
			}
			newimg.Set(x, y, newpix)

		}
	}
	return newimg
}
func GetRed(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newimg := image.NewRGBA(bounds)

	w, h := bounds.Max.X, bounds.Max.Y

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := img.At(x, y)
			r, _, _, a := pix.RGBA()

			newpix := color.RGBA{
				R: uint8(r),
				G: uint8(0),
				B: uint8(0),
				A: uint8(a),
			}
			newimg.Set(x, y, newpix)

		}
	}
	return newimg
}

func GrayColors(img image.Image) *image.Gray {
	bounds := img.Bounds()

	grayimg := image.NewGray(bounds)

	w, h := bounds.Max.X, bounds.Max.Y

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := img.At(x, y)

			grayc := color.GrayModel.Convert(pix)

			grayimg.Set(x, y, grayc)
		}
	}
	return grayimg
}
func AsciImage(img image.Gray) ([]string, error) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var toret []string

	for y := 0; y < h; y++ {
		s := ""
		for x := 0; x < w; x++ {
			pix := img.At(x, y)
			grayColor, ok := pix.(color.Gray)
			if !ok {
				return nil, fmt.Errorf("image is not gray")
			}

			val := grayColor.Y
			val = val / (255/uint8(len(LETTERS)) + 1)

			s += LETTERS[val]

		}
		toret = append(toret, s)
	}
	return toret, nil

}

func ResizeImage(img image.Image, factor int) *image.RGBA {
	bounds := img.Bounds()
	w, h := bounds.Max.X/factor, bounds.Max.Y/factor
	newRect := image.Rect(0, 0, w, h)

	new := image.NewRGBA(newRect)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := img.At(x*factor, y*factor)

			new.Set(x, y, pix)
		}
	}
	return new
}
