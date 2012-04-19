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
	"add":           blend.Add,
	"color":         blend.Color,
	"color_burn":    blend.ColorBurn,
	"color_dodge":   blend.ColorDodge,
	"darken":        blend.Darken,
	"darker_color":  blend.DarkerColor,
	"difference":    blend.Difference,
	"divide":        blend.Divide,
	"exclusion":     blend.Exclusion,
	"hard_light":    blend.HardLight,
	"hard_mix":      blend.HardMix,
	"hue":           blend.Hue,
	"lighten":       blend.Lighten,
	"lighter_color": blend.LighterColor,
	"linear_burn":   blend.LinearBurn,
	"linear_dodge":  blend.LinearDodge,
	"linear_light":  blend.LinearLight,
	"luminosity":    blend.Luminosity,
	"multiply":      blend.Multiply,
	"overlay":       blend.Overlay,
	"phoenix":       blend.Phoenix,
	"pin_light":     blend.PinLight,
	"reflex":        blend.Reflex,
	"saturation":    blend.Saturation,
	"screen":        blend.Screen,
	"soft_light":    blend.SoftLight,
	"substract":     blend.Substract,
	"vivid_light":   blend.VividLight,
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
		err = SaveJPG("blend_"+name, img)
		if err != nil {
			panic(err)
		}
	}
}
