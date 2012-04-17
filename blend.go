package blend

import (
	"image/color"
	"math"
)

const (
	max = 65535.0 // equals to 0xFFFF uint16 max range of color.Color
	mid = max / 2.0
)

type BlendFunc func(src color.Color, dst color.Color) color.Color

func blend_per_channel(src, dst color.Color, bf func(float64, float64) float64) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	return rgbaf64{bf(s.r, d.r), bf(s.g, d.g), bf(s.b, d.b), d.a}
}

// ADD
func ADD(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, add)
}
func add(s, d float64) float64 {
	if s+d > max {
		return max
	}
	return s + d
}

// SUBSTRACT
func SUBSTRACT(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, substract)
}
func substract(s, d float64) float64 {
	if d-s < 0.0 {
		return 0.0
	}
	return d - s
}

// DIVIDE
func DIVIDE(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, divide)
}
func divide(s, d float64) float64 {
	return d / s
}

// MULTIPLY
func MULTIPLY(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, multiply)
}
func multiply(s, d float64) float64 {
	return s * d / max
}

// SCREEN
func SCREEN(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, screen)
}
func screen(s, d float64) float64 {
	return s + d - s*d/max
}

// OVERLAY
func OVERLAY(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, overlay)
}
func overlay(s, d float64) float64 {
	if d < mid {
		return 2 * s * d / max
	}
	return max - 2*(max-s)*(max-d)/max
}

// SOFT LIGHT
func SOFT_LIGHT(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, soft_light)
}
func soft_light(s, d float64) float64 {
	if s > mid {
		return d + (max-d)*((s-mid)/mid)*(0.5-math.Abs(d-mid)/max)
	}
	return d - d*((mid-s)/mid)*(0.5-math.Abs(d-mid)/max)
}

// HARD LIGHT
func HARD_LIGHT(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, hard_light)
}
func hard_light(s, d float64) float64 {
	if s > mid {
		return d + (max-d)*((s-mid)/mid)
	}
	return d * s / mid
}

// COLOR DODGE
func COLOR_DODGE(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, color_dodge)
}
func color_dodge(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * max / (max - s)))
}

// COLOR BURN
func COLOR_BURN(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, color_burn)
}
func color_burn(s, d float64) float64 {
	if s == 0.0 {
		return s
	}
	return math.Max(0.0, max-((max-d)*max/s))
}

// LINEAR DODGE
func LINEAR_DODGE(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_dodge)
}
func linear_dodge(s, d float64) float64 {
	return math.Min(s+d, max)
}

// LINEAR BURN
func LINEAR_BURN(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_burn)
}
func linear_burn(s, d float64) float64 {
	if (s + d) < max {
		return 0.0
	}
	return s + d - max
}

// DARKEN
func DARKEN(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, darken)
}
func darken(s, d float64) float64 {
	return math.Min(s, d)
}

// LIGHTEN
func LIGHTEN(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, lighten)
}
func lighten(s, d float64) float64 {
	return math.Max(s, d)
}

// DIFFERENCE
func DIFFERENCE(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, difference)
}
func difference(s, d float64) float64 {
	return math.Abs(s - d)
}

// EXCLUSION
func EXCLUSION(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, exclusion)
}
func exclusion(s, d float64) float64 {
	return s + d - s*d/mid
}

// REFLEX
func REFLEX(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, reflex)
}
func reflex(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * d / (max - s)))
}

// LINEAR LIGHT
func LINEAR_LIGHT(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_light)
}
func linear_light(s, d float64) float64 {
	if s < mid {
		return linear_burn(d, (2 * s))
	}
	return linear_dodge(d, (2 * (s - mid)))
}

// PIN LIGHT
func PIN_LIGHT(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, pin_light)
}
func pin_light(s, d float64) float64 {
	if s < mid {
		return darken(d, (2 * s))
	}
	return lighten(d, (2 * (s - mid)))
}

// VIVID LIGHT
func VIVID_LIGHT(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, vivid_light)
}
func vivid_light(s, d float64) float64 {
	if s < mid {
		return color_burn(d, (2 * s))
	}
	return color_dodge(d, (2 * (s - mid)))
}

// HARD MIX
func HARD_MIX(src color.Color, dst color.Color) color.Color {
	return blend_per_channel(src, dst, hard_mix)
}
func hard_mix(s, d float64) float64 {
	if vivid_light(s, d) < mid {
		return 0.0
	}
	return max
}
