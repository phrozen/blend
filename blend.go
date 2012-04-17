package blend

import (
	"image/color"
	"math"
)

const (
	MULTIPLY = iota
	SCREEN
	OVERLAY
	SOFT_LIGHT
	HARD_LIGHT
	COLOR_DODGE
	COLOR_BURN
	LINEAR_COLOR_DODGE
	LINEAR_COLOR_BURN
	DARKEN
	LIGHTEN
	DIFFERENCE
	EXCLUSION
	REFLEX
	LINEAR_LIGHT
	PIN_LIGHT
	VIVID_LIGHT
	HARD_MIX
	// Blending modes in HSL color model.
	HUE
	COLOR
	LUMINOSITY
)

const (
	max = 65535.0
	mid = max / 2.0
)

type blendFunc func(float64, float64) float64

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

func colorTorgbaf64(c color.Color) rgbaf64 {
	r, g, b, a := c.RGBA()
	return rgbaf64{float64(r), float64(g), float64(b), float64(a)}
}

func BlendColor(source, dest color.Color, mode int) color.Color {
	switch mode {
	case MULTIPLY:
		return blend(source, dest, multiply)
	case SCREEN:
		return blend(source, dest, screen)
	case OVERLAY:
		return blend(source, dest, overlay)
	case SOFT_LIGHT:
		return blend(source, dest, soft_light)
	case HARD_LIGHT:
		return blend(source, dest, hard_light)
	}
	return rgbaf64{0.0, 0.0, 0.0, 0.0}
}

func blend(source, dest color.Color, bf blendFunc) (c rgbaf64) {
	s := colorTorgbaf64(source)
	d := colorTorgbaf64(dest)
	c.r = bf(s.r, d.r)
	c.r = bf(s.g, d.g)
	c.b = bf(s.b, d.b)
	c.a = s.a
	return
}

// Blend Modes
func multiply(s, d float64) float64 {
	return s * d / max
}

func screen(s, d float64) float64 {
	return s + d - s*d/max
}

func overlay(s, d float64) float64 {
	if d < mid {
		return 2 * s * d / max
	}
	return max - 2*(max-s)*(max-d)/max
}

func soft_light(s, d float64) float64 {
	if s > mid {
		return d + (max-d)*((s-mid)/mid)*(0.5-math.Abs(d-mid)/max)
	}
	return d - d*((mid-s)/mid)*(0.5-math.Abs(d-mid)/max)
}

func hard_light(s, d float64) float64 {
	if s > mid {
		return d + (max-d)*((s-mid)/mid)
	}
	return d * s / mid
}
