package blend

import (
	"image/color"
	"math"
)

var (
	MULTIPLY     = multiply
	SCREEN       = screen
	OVERLAY      = overlay
	SOFT_LIGHT   = soft_light
	HARD_LIGHT   = hard_light
	COLOR_DODGE  = color_dodge
	COLOR_BURN   = color_burn
	LINEAR_DODGE = linear_dodge
	LINEAR_BURN  = linear_burn
	DARKEN       = darken
	LIGHTEN      = lighten
	DIFFERENCE   = difference
	EXCLUSION    = exclusion
	REFLEX       = reflex
	LINEAR_LIGHT = linear_light
	PIN_LIGHT    = pin_light
	VIVID_LIGHT  = vivid_light
	HARD_MIX     = hard_mix
	// Blending modes in HSL color model.
	/*
		HUE
		COLOR
		LUMINOSITY
	*/
)

const (
	max = 65535.0
	mid = max / 2.0
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

func colorTorgbaf64(c color.Color) rgbaf64 {
	r, g, b, a := c.RGBA()
	return rgbaf64{float64(r), float64(g), float64(b), float64(a)}
}

type BlendFunc func(float64, float64) float64

func Blend(src, dst color.Color, bf BlendFunc) color.Color {
	s := colorTorgbaf64(src)
	d := colorTorgbaf64(dst)
	return rgbaf64{bf(s.r, d.r), bf(s.g, d.g), bf(s.b, d.b), d.a}
}

// Blend Functions
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

func color_dodge(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * max / (max - s)))
}

func color_burn(s, d float64) float64 {
	if s == 0.0 {
		return s
	}
	return math.Max(0.0, max-((max-d)*max/s))
}

func linear_dodge(s, d float64) float64 {
	return math.Min(s+d, max)
}

func linear_burn(s, d float64) float64 {
	if (s + d) < max {
		return 0.0
	}
	return s + d - max
}

func darken(s, d float64) float64 {
	return math.Min(s, d)
}

func lighten(s, d float64) float64 {
	return math.Max(s, d)
}

func difference(s, d float64) float64 {
	return math.Abs(s - d)
}

func exclusion(s, d float64) float64 {
	return s + d - s*d/mid
}

func reflex(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * d / (max - s)))
}

func linear_light(s, d float64) float64 {
	if s < mid {
		return linear_burn(d, (2 * s))
	}
	return linear_dodge(d, (2 * (s - mid)))
}

func pin_light(s, d float64) float64 {
	if s < mid {
		return darken(d, (2 * s))
	}
	return lighten(d, (2 * (s - mid)))
}

func vivid_light(s, d float64) float64 {
	if s < mid {
		return color_burn(d, (2 * s))
	}
	return color_dodge(d, (2 * (s - mid)))
}

func hard_mix(s, d float64) float64 {
	if vivid_light(s, d) < mid {
		return 0.0
	}
	return max
}
