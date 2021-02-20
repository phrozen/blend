/*
	Example of blending modes using blend Go library.
	https://github.com/phrozen/blend
	by Guillermo Estrada
*/

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/phrozen/blend"
)

func loadJPG(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveJPG(filename string, img image.Image) error {
	file, err := os.Create(filename + ".jpg")
	if err != nil {
		return err
	}

	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 85}); err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var err error
	var img image.Image

	dst, err := loadJPG("destination.jpg")
	if err != nil {
		panic(err)
	}

	src, err := loadJPG("source.jpg")
	if err != nil {
		panic(err)
	}

	fmt.Println("This program tests all the color blending modes in the library.")

	for name, mode := range blend.Modes {
		fmt.Println("Blending Mode: ", name)
		img = blend.BlendNewImage(dst, src, mode)
		err = saveJPG("output/blend_"+name, img)
		if err != nil {
			panic(err)
		}
	}
}
