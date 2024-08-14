package myimage

import (
	"image"
	"image/color"
)

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
