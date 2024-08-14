package uitilities

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func GetFile(str string) (image.Image, error) {
	path := "./images/" + str
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return img, nil

}
