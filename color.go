
package blend

import (
	"image/color"
)

type rgbaf64 struct {
	r, g, b, a float64
}

func (c rgbaf64) RGBA() (uint32, uint32, uint32, uint32) {
	r := float64ToUint16(c.r)
	g := float64ToUint16(c.g)
	b := float64ToUint16(c.b)
	a := float64ToUint16(c.a)
	return uint32(r), uint32(g), uint32(b), uint32(a)
}

func color2rgbaf64(c color.Color) rgbaf64 {
	r, g, b, a := c.RGBA()
	return rgbaf64{float64(r), float64(g), float64(b), float64(a)}
}

func HSLtoRGB() {
	
}

func RGBtoHSL() {
	
}
