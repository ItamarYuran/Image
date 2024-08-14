package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func GetFile(str string) (image.Image, error) {
	parts := strings.Split(str, ".")
	surffix := parts[1]

	path := "../../images/" + str
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var img image.Image

	switch surffix {

	case "jpeg", "jpg":
		img, err = jpeg.Decode(file)

	case "png":
		img, err = png.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", surffix)
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return img, nil

}

func SaveFile(name string, img image.Image) {
	parts := strings.Split(name, ".")
	surffix := parts[1]

	path := "../../images/" + name
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()
	switch surffix {

	case "jpeg", "jpg":
		err = jpeg.Encode(out, img, nil)

	case "png":
		err = png.Encode(out, img)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

}
