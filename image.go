package blend

import (
	"image"
)

func BlendImage(source, dest image.Image, mode BlendFunc) (image.Image, error) {

	if source.ColorModel() != dest.ColorModel() {
		return nil, BlendError{"Source and destination images have different color models."}
	}

	swidth, sheight := getWidthAndHeight(source)
	dwidth, dheight := getWidthAndHeight(dest)

	if swidth != dwidth {
		return nil, BlendError{"Source and destination images have different widths."}
	}
	if sheight != dheight {
		return nil, BlendError{"Source and destination images have different heights."}
	}

	img := image.NewRGBA64(image.Rect(0, 0, dwidth, dheight))

	for x := 0; x < dwidth; x++ {
		for y := 0; y < dheight; y++ {
			img.Set(x, y, Blend(source.At(x, y), dest.At(x, y), mode))
		}
	}

	return img, nil
}

func getWidthAndHeight(img image.Image) (width, height int) {
	width = img.Bounds().Max.X - img.Bounds().Min.X
	height = img.Bounds().Max.Y - img.Bounds().Min.Y
	return
}
