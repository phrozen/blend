/*
	Example of blending modes using blend Go library.
	https://github.com/phrozen/blend
	by Guillermo Estrada
*/

package main

import (
	"fmt"
	"github.com/phrozen/blend"
	"image"
	"image/jpeg"
	"os"
)

var modes = map[string]blend.BlendFunc{
	"add":           blend.ADD,
	"substract":     blend.SUBSTRACT,
	"divide":        blend.DIVIDE,
	"multiply":      blend.MULTIPLY,
	"screen":        blend.SCREEN,
	"overlay":       blend.OVERLAY,
	"soft_light":    blend.SOFT_LIGHT,
	"hard_light":    blend.HARD_LIGHT,
	"color_dodge":   blend.COLOR_DODGE,
	"color_burn":    blend.COLOR_BURN,
	"linear_dodge":  blend.LINEAR_DODGE,
	"linear_burn":   blend.LINEAR_BURN,
	"darken":        blend.DARKEN,
	"lighten":       blend.LIGHTEN,
	"difference":    blend.DIFFERENCE,
	"exclusion":     blend.EXCLUSION,
	"reflex":        blend.REFLEX,
	"linear_light":  blend.LINEAR_LIGHT,
	"pin_light":     blend.PIN_LIGHT,
	"vivid_light":   blend.VIVID_LIGHT,
	"hard_mix":      blend.HARD_MIX,
	"darker_color":  blend.DARKER_COLOR,
	"lighter_color": blend.LIGHTER_COLOR,
	"phoenix":       blend.PHOENIX,
	"hue":           blend.HUE,
	"saturation":    blend.SATURATION,
	"color":         blend.COLOR,
	"luminosity":    blend.LUMINOSITY,
}

func LoadJPG(filename string) (image.Image, error) {
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

func SaveJPG(filename string, img image.Image) error {
	file, err := os.Create(filename + ".jpg")
	if err != nil {
		return err
	}

	if err := jpeg.Encode(file, img, &jpeg.Options{85}); err != nil {
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

	dst, err := LoadJPG("destination.jpg")
	if err != nil {
		panic(err)
	}

	src, err := LoadJPG("source.jpg")
	if err != nil {
		panic(err)
	}

	fmt.Println("This program tests all the color blending modes in the library.")

	for name, mode := range modes {
		fmt.Println("Blending Mode: ", name)
		img, err = blend.Blend(dst, src, mode)
		if err != nil {
			panic(err)
		}
		err = SaveJPG(name, img)
		if err != nil {
			panic(err)
		}
	}
}
