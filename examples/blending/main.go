package main

import (
	"fmt"
	"github.com/Phrozen/blend"
	"image"
	"image/png"
	"os"
)

var modes1 = map[string]blend.BlendFunc{
	"multiply":     blend.MULTIPLY,
	"screen":       blend.SCREEN,
	"overlay":      blend.OVERLAY,
	"soft_light":   blend.SOFT_LIGHT,
	"hard_light":   blend.HARD_LIGHT,
	"color_dodge":  blend.COLOR_DODGE,
	"color_burn":   blend.COLOR_BURN,
	"linear_dodge": blend.LINEAR_DODGE,
	"linear_burn":  blend.LINEAR_BURN,
	"darken":       blend.DARKEN,
	"lighten":      blend.LIGHTEN,
	"difference":   blend.DIFFERENCE,
	"exclusion":    blend.EXCLUSION,
	"reflex":       blend.REFLEX,
}

var modes2 = map[string]blend.BlendFunc{
	"linear_light": blend.LINEAR_LIGHT,
	"pin_light":    blend.PIN_LIGHT,
	"vivid_light":  blend.VIVID_LIGHT,
	"hard_mix":     blend.HARD_MIX,
}

func LoadPNG(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SavePNG(filename string, img image.Image) error {
	file, err := os.Create(filename + ".png")
	if err != nil {
		return err
	}

	if err := png.Encode(file, img); err != nil {
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

	dst, err := LoadPNG("dest.png")
	if err != nil {
		panic(err)
	}

	src1, err := LoadPNG("source1.png")
	if err != nil {
		panic(err)
	}

	src2, err := LoadPNG("source2.png")
	if err != nil {
		panic(err)
	}

	fmt.Println("This program tests all the color blending modes in the library.")

	// Testing Blending Modes with source1.png
	for key, value := range modes1 {
		fmt.Print("*Blending Mode: ", key, " ...")
		img, err = blend.BlendImage(src1, dst, value)
		if err != nil {
			panic(err)
		}
		err = SavePNG(key, img)
		if err != nil {
			panic(err)
		}
		fmt.Println("Saved! ", key, ".png")
	}

	// Testing Blending Modes with source2.png
	for key, value := range modes2 {
		fmt.Print("*Blending Mode: ", key, " ...")
		img, err = blend.BlendImage(src2, dst, value)
		if err != nil {
			panic(err)
		}
		err = SavePNG(key, img)
		if err != nil {
			panic(err)
		}
		fmt.Println("Saved! ", key, ".png")
	}

}
