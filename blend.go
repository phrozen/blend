// Copyright (c) 2012 Guillermo Estrada. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package image implements blending mode functions bewteen images.
//
// The fundamental part of the library id the type BlendFunc,
// the function is applied to each pixel where the top layer (src)
// overlaps the bottom layer (dst) of both given 'image' interfaces.
//
// This library provides a lot of Blend Functions to be used either
// as 'mode' parameter to the Blend() primary function, or to use 
// individually providing two 'color' interfaces.
//
// See documentation for more details:
// http://github.com/phrozen/blend
package blend

import (
	"image"
	"image/color"
	"math"
)

// Constants of max and mid values for uint16 for internal use.
// This can be changed to make the algorithms use uint8 instead, 
// but they are kept this way to provide more acurate calculations
// and to support all of the color modes in the 'image' package.
const (
	max = 65535.0 // equals to 0xFFFF uint16 max range of color.Color
	mid = max / 2.0
)

// Blends src image (top layer) into dst image (bottom layer) using
// the BlendFunc provided by mode. BlendFunc is applied to each pixel
// where the src image overlaps the dst image and returns the resulting
// image or an error in case of failure.
func Blend(src, dst image.Image, mode BlendFunc) (image.Image, error) {

	// Color model check. Needs more testing to see if there is no problem 
	// using the interfaces, to blend images with different color models.
	if src.ColorModel() != dst.ColorModel() {
		return nil, BlendError{"Top layer(src) and bot layer(dst) have different color models."}
	}

	// Boundary check to see if we can blend all pixels in the top layer
	// into the bottom layer. Later and intersection will be used.
	if !src.Bounds().In(dst.Bounds()) {
		return nil, BlendError{"Top layer(src) does not fit into bottom layer(dst)."}
	}

	// Create a new RGBA or RGBA64 image to return the values.
	img := image.NewRGBA(dst.Bounds())

	for x := 0; x < dst.Bounds().Dx(); x++ {
		for y := 0; y < dst.Bounds().Dy(); y++ {
			// If src is inside dst, we blend both pixels
			if p := image.Pt(x, y); p.In(src.Bounds()) {
				img.Set(x, y, mode(src.At(x, y), dst.At(x, y)))
			} else {
				// else we copy dst pixel.
				img.Set(x, y, dst.At(x, y))
			}
		}
	}
	return img, nil
}

type BlendFunc func(src, dst color.Color) color.Color

func blend_per_channel(src, dst color.Color, bf func(float64, float64) float64) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	return rgbaf64{bf(s.r, d.r), bf(s.g, d.g), bf(s.b, d.b), d.a}
}

// Blending modes supported by Photoshop in order.
/*-------------------------------------------------------*/

// DARKEN
func DARKEN(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, darken)
}
func darken(s, d float64) float64 {
	return math.Min(s, d)
}

// MULTIPLY
func MULTIPLY(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, multiply)
}
func multiply(s, d float64) float64 {
	return s * d / max
}

// COLOR BURN
func COLOR_BURN(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, color_burn)
}
func color_burn(s, d float64) float64 {
	if s == 0.0 {
		return s
	}
	return math.Max(0.0, max-((max-d)*max/s))
}

// LINEAR BURN
func LINEAR_BURN(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_burn)
}
func linear_burn(s, d float64) float64 {
	if (s + d) < max {
		return 0.0
	}
	return s + d - max
}

// DARKER COLOR
func DARKER_COLOR(src, dst color.Color) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	if s.r+s.g+s.b > d.r+d.g+d.b {
		return dst
	}
	return src
}

/*-------------------------------------------------------*/

// LIGHTEN
func LIGHTEN(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, lighten)
}
func lighten(s, d float64) float64 {
	return math.Max(s, d)
}

// SCREEN
func SCREEN(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, screen)
}
func screen(s, d float64) float64 {
	return s + d - s*d/max
}

// COLOR DODGE
func COLOR_DODGE(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, color_dodge)
}
func color_dodge(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * max / (max - s)))
}

// LINEAR DODGE
func LINEAR_DODGE(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_dodge)
}
func linear_dodge(s, d float64) float64 {
	return math.Min(s+d, max)
}

// LIGHTER COLOR
func LIGHTER_COLOR(src, dst color.Color) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	if s.r+s.g+s.b > d.r+d.g+d.b {
		return src
	}
	return dst
}

/*-------------------------------------------------------*/

// OVERLAY
func OVERLAY(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, overlay)
}
func overlay(s, d float64) float64 {
	if d < mid {
		return 2 * s * d / max
	}
	return max - 2*(max-s)*(max-d)/max
}

// SOFT LIGHT
func SOFT_LIGHT(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, soft_light)
}
func soft_light(s, d float64) float64 {
	return (d / max) * (d + (2*s/max)*(max-d))
}

// HARD LIGHT
func HARD_LIGHT(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, hard_light)
}
func hard_light(s, d float64) float64 {
	if s > mid {
		return d + (max-d)*((s-mid)/mid)
	}
	return d * s / mid
}

// VIVID LIGHT (check)
func VIVID_LIGHT(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, vivid_light)
}
func vivid_light(s, d float64) float64 {
	if s < mid {
		return color_burn(d, (2 * s))
	}
	return color_dodge(d, (2 * (s - mid)))
}

// LINEAR LIGHT
func LINEAR_LIGHT(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_light)
}
func linear_light(s, d float64) float64 {
	if s < mid {
		return linear_burn(d, (2 * s))
	}
	return linear_dodge(d, (2 * (s - mid)))
}

// PIN LIGHT
func PIN_LIGHT(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, pin_light)
}
func pin_light(s, d float64) float64 {
	if s < mid {
		return darken(d, (2 * s))
	}
	return lighten(d, (2 * (s - mid)))
}

// HARD MIX (check)
func HARD_MIX(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, hard_mix)
}
func hard_mix(s, d float64) float64 {
	if vivid_light(s, d) < mid {
		return 0.0
	}
	return max
}

/*-------------------------------------------------------*/

// DIFFERENCE
func DIFFERENCE(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, difference)
}
func difference(s, d float64) float64 {
	return math.Abs(s - d)
}

// EXCLUSION
func EXCLUSION(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, exclusion)
}
func exclusion(s, d float64) float64 {
	return s + d - s*d/mid
}

// SUBSTRACT
func SUBSTRACT(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, substract)
}
func substract(s, d float64) float64 {
	if d-s < 0.0 {
		return 0.0
	}
	return d - s
}

// DIVIDE
func DIVIDE(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, divide)
}
func divide(s, d float64) float64 {
	return (d*max)/s + 1.0
}

// Blending modes that use HSL color model transformations.
/*-------------------------------------------------------*/

// HUE
func HUE(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	if s.s == 0.0 {
		return dst
	}
	d := rgb2hsl(dst)
	return hsl2rgb(s.h, d.s, d.l)
}

// SATURATION
func SATURATION(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(d.h, s.s, d.l)
}

// COLOR
func COLOR(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(s.h, s.s, d.l)
}

// LUMINOSITY
func LUMINOSITY(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(d.h, d.s, s.l)
}

// This blending modes are not implemented in Photoshop
// or GIMP at the moment, but produced their desired results.
/*-------------------------------------------------------*/

// ADD
func ADD(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, add)
}
func add(s, d float64) float64 {
	if s+d > max {
		return max
	}
	return s + d
}

// REFLEX (a.k.a GLOW)
func REFLEX(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, reflex)
}
func reflex(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * d / (max - s)))
}

// PHOENIX
func PHOENIX(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, phoenix)
}
func phoenix(s, d float64) float64 {
	return math.Min(s, d) - math.Max(s, d) + max
}
