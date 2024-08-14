package main

import (
	"fmt"

	utils "github.com/ItamarYuran/MyImage/pkg/imageproc"
)

func main() {
	img, err := utils.GetFile("meandsarah.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	img2 := utils.ResizeImage(img, 17)
	utils.SaveFile("resized.jpeg", img2)

	imgray := utils.GrayColors(img2)

	arr, _ := utils.AsciImage(*imgray)

	for _, a := range arr {
		fmt.Println(a)
	}

}
