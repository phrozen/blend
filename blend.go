// Copyright (c) 2012 Guillermo Estrada. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package blend implements blending mode functions bewteen images,
// and some utility functions for image processing.
//
// The fundamental part of the library is the type BlendFunc,
// the function is applied to each pixel where the top layer (src)
// overlaps the bottom layer (dst) of both given 'image' interfaces.
//
// This library provides many of the widely used blending functions 
// to be used either as 'mode' parameter to the Blend() primary
// function, or to be used individually providing two 'color' interfaces.
// You can implement your own blending modes and pass them into the 
// Blend() function.
//
// This is the list of the currently implemented blending modes:
//
// Add, Color, Color Burn, Color Dodge, Darken, Darker Color, Difference, 
// Divide, Exclusion, Hard Light, Hard Mix, Hue, Lighten, Lighter Color, 
// Linear Burn, Linear Dodge, Linear Light, Luminosity, Multiply, Overlay, 
// Phoenix, Pin Light, Reflex, Saturation, Screen, Soft Light, Substract, 
// Vivid Light.
//
// Check github for more details:
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

var (
	ADD           BlendFunc
	COLOR         BlendFunc
	COLOR_BURN    BlendFunc
	COLOR_DODGE   BlendFunc
	DARKEN        BlendFunc
	DARKER_COLOR  BlendFunc
	DIFFERENCE    BlendFunc
	DIVIDE        BlendFunc
	EXCLUSION     BlendFunc
	HARD_LIGHT    BlendFunc
	HARD_MIX      BlendFunc
	HUE           BlendFunc
	LIGHTEN       BlendFunc
	LIGHTER_COLOR BlendFunc
	LINEAR_BURN   BlendFunc
	LINEAR_DODGE  BlendFunc
	LINEAR_LIGHT  BlendFunc
	LUMINOSITY    BlendFunc
	MULTIPLY      BlendFunc
	OVERLAY       BlendFunc
	PHOENIX       BlendFunc
	PIN_LIGHT     BlendFunc
	REFLEX        BlendFunc
	SATURATION    BlendFunc
	SCREEN        BlendFunc
	SOFT_LIGHT    BlendFunc
	SUBSTRACT     BlendFunc
	VIVID_LIGHT   BlendFunc
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
func darken(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, darken_per_ch)
}
func darken_per_ch(s, d float64) float64 {
	return math.Min(s, d)
}

// MULTIPLY
func multiply(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, multiply_per_ch)
}
func multiply_per_ch(s, d float64) float64 {
	return s * d / max
}

// COLOR BURN
func color_burn(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, color_burn_per_ch)
}
func color_burn_per_ch(s, d float64) float64 {
	if s == 0.0 {
		return s
	}
	return math.Max(0.0, max-((max-d)*max/s))
}

// LINEAR BURN
func linear_burn(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_burn_per_ch)
}
func linear_burn_per_ch(s, d float64) float64 {
	if (s + d) < max {
		return 0.0
	}
	return s + d - max
}

// DARKER COLOR
func darker_color(src, dst color.Color) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	if s.r+s.g+s.b > d.r+d.g+d.b {
		return dst
	}
	return src
}

/*-------------------------------------------------------*/

// LIGHTEN
func lighten(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, lighten_per_ch)
}
func lighten_per_ch(s, d float64) float64 {
	return math.Max(s, d)
}

// SCREEN
func screen(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, screen_per_ch)
}
func screen_per_ch(s, d float64) float64 {
	return s + d - s*d/max
}

// COLOR DODGE
func color_dodge(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, color_dodge_per_ch)
}
func color_dodge_per_ch(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * max / (max - s)))
}

// LINEAR DODGE
func linear_dodge(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_dodge_per_ch)
}
func linear_dodge_per_ch(s, d float64) float64 {
	return math.Min(s+d, max)
}

// LIGHTER COLOR
func lighter_color(src, dst color.Color) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	if s.r+s.g+s.b > d.r+d.g+d.b {
		return src
	}
	return dst
}

/*-------------------------------------------------------*/

// OVERLAY
func overlay(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, overlay_per_ch)
}
func overlay_per_ch(s, d float64) float64 {
	if d < mid {
		return 2 * s * d / max
	}
	return max - 2*(max-s)*(max-d)/max
}

// SOFT LIGHT
func soft_light(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, soft_light_per_ch)
}
func soft_light_per_ch(s, d float64) float64 {
	return (d / max) * (d + (2*s/max)*(max-d))
}

// HARD LIGHT
func hard_light(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, hard_light_per_ch)
}
func hard_light_per_ch(s, d float64) float64 {
	if s > mid {
		return d + (max-d)*((s-mid)/mid)
	}
	return d * s / mid
}

// VIVID LIGHT (check)
func vivid_light(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, vivid_light_per_ch)
}
func vivid_light_per_ch(s, d float64) float64 {
	if s < mid {
		return color_burn_per_ch(d, (2 * s))
	}
	return color_dodge_per_ch(d, (2 * (s - mid)))
}

// LINEAR LIGHT
func linear_light(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, linear_light_per_ch)
}
func linear_light_per_ch(s, d float64) float64 {
	if s < mid {
		return linear_burn_per_ch(d, (2 * s))
	}
	return linear_dodge_per_ch(d, (2 * (s - mid)))
}

// PIN LIGHT
func pin_light(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, pin_light_per_ch)
}
func pin_light_per_ch(s, d float64) float64 {
	if s < mid {
		return darken_per_ch(d, (2 * s))
	}
	return lighten_per_ch(d, (2 * (s - mid)))
}

// HARD MIX (check)
func hard_mix(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, hard_mix_per_ch)
}
func hard_mix_per_ch(s, d float64) float64 {
	if vivid_light_per_ch(s, d) < mid {
		return 0.0
	}
	return max
}

/*-------------------------------------------------------*/

// DIFFERENCE
func difference(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, difference_per_ch)
}
func difference_per_ch(s, d float64) float64 {
	return math.Abs(s - d)
}

// EXCLUSION
func exclusion(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, exclusion_per_ch)
}
func exclusion_per_ch(s, d float64) float64 {
	return s + d - s*d/mid
}

// SUBSTRACT
func substract(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, substract_per_ch)
}
func substract_per_ch(s, d float64) float64 {
	if d-s < 0.0 {
		return 0.0
	}
	return d - s
}

// DIVIDE
func divide(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, divide_per_ch)
}
func divide_per_ch(s, d float64) float64 {
	return (d*max)/s + 1.0
}

// Blending modes that use HSL color model transformations.
/*-------------------------------------------------------*/

// HUE
func hue(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	if s.s == 0.0 {
		return dst
	}
	d := rgb2hsl(dst)
	return hsl2rgb(s.h, d.s, d.l)
}

// SATURATION
func saturation(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(d.h, s.s, d.l)
}

// COLOR "added _ to avoid namespace conflict with 'color' package"
func color_(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(s.h, s.s, d.l)
}

// LUMINOSITY
func luminosity(src, dst color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(d.h, d.s, s.l)
}

// This blending modes are not implemented in Photoshop
// or GIMP at the moment, but produced their desired results.
/*-------------------------------------------------------*/

// ADD
func add(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, add_per_ch)
}
func add_per_ch(s, d float64) float64 {
	if s+d > max {
		return max
	}
	return s + d
}

// REFLEX (a.k.a GLOW)
func reflex(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, reflex_per_ch)
}
func reflex_per_ch(s, d float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * d / (max - s)))
}

// PHOENIX
func phoenix(src, dst color.Color) color.Color {
	return blend_per_channel(src, dst, phoenix_per_ch)
}
func phoenix_per_ch(s, d float64) float64 {
	return math.Min(s, d) - math.Max(s, d) + max
}

// Init function maps the blendingmode functions.
func init() {
	DARKEN = darken
	MULTIPLY = multiply
	COLOR_BURN = color_burn
	LINEAR_BURN = linear_burn
	DARKER_COLOR = darker_color
	LIGHTEN = lighten
	SCREEN = screen
	COLOR_DODGE = color_dodge
	LINEAR_DODGE = linear_dodge
	LIGHTER_COLOR = lighter_color
	OVERLAY = overlay
	SOFT_LIGHT = soft_light
	HARD_LIGHT = hard_light
	VIVID_LIGHT = vivid_light
	LINEAR_LIGHT = linear_light
	PIN_LIGHT = pin_light
	HARD_MIX = hard_mix
	DIFFERENCE = difference
	EXCLUSION = exclusion
	SUBSTRACT = substract
	DIVIDE = divide
	HUE = hue
	SATURATION = saturation
	COLOR = color_
	LUMINOSITY = luminosity
	ADD = add
	REFLEX = reflex
	PHOENIX = phoenix
}
