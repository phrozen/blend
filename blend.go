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
	"image/draw"
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

// BlendFunc or blend mode receives a destination color and
// a source color, then returns a transformation of them. Blend()
// function receives a BlendFunc and applies it to every pixel in
// the overlaping areas of two given images.
type BlendFunc func(dst, src color.Color) color.Color

// Available Moodes map
var Modes = map[string]BlendFunc{
	"add":           Add,
	"color":         Color,
	"color_burn":    ColorBurn,
	"color_dodge":   ColorDodge,
	"darken":        Darken,
	"darker_color":  DarkerColor,
	"difference":    Difference,
	"divide":        Divide,
	"exclusion":     Exclusion,
	"hard_light":    HardLight,
	"hard_mix":      HardMix,
	"hue":           Hue,
	"lighten":       Lighten,
	"lighter_color": LighterColor,
	"linear_burn":   LinearBurn,
	"linear_dodge":  LinearDodge,
	"linear_light":  LinearLight,
	"luminosity":    Luminosity,
	"multiply":      Multiply,
	"overlay":       Overlay,
	"phoenix":       Phoenix,
	"pin_light":     PinLight,
	"reflex":        Reflex,
	"saturation":    Saturation,
	"screen":        Screen,
	"soft_light":    SoftLight,
	"substract":     Substract,
	"vivid_light":   VividLight,
}

// BlendImage blends src image (top layer) into dst image (bottom layer) using
// the BlendFunc provided by mode. BlendFunc is applied to each pixel
// where the src image overlaps the dst image and the result is stored
// in the original dst image, src image is unmutable.
func BlendImage(dst draw.Image, src image.Image, mode BlendFunc) {
	// Obtain the intersection of both images.
	inter := dst.Bounds().Intersect(src.Bounds())
	// Apply BlendFuc to each pixel in the intersection.
	for y := inter.Min.Y; y < inter.Max.Y; y++ {
		for x := inter.Min.X; x < inter.Max.X; x++ {
			dst.Set(x, y, mode(dst.At(x, y), src.At(x, y)))
		}
	}
}

// BlendNewImage blends src image (top layer) into dst image (bottom layer) using
// the BlendFunc provided by mode. BlendFunc is applied to each pixel
// where the src image overlaps the dst image and returns the resulting
// image without modifying src, or dst as they are both unmutable.
func BlendNewImage(dst, src image.Image, mode BlendFunc) image.Image {
	// Obtain the intersection of both images.
	inter := dst.Bounds().Intersect(src.Bounds())
	// Create a new RGBA or RGBA64 image to return the values.
	img := image.NewRGBA(dst.Bounds())
	// Iterate over dst image pixels.
	for y := dst.Bounds().Min.Y; y < dst.Bounds().Max.Y; y++ {
		for x := dst.Bounds().Min.X; x < dst.Bounds().Max.X; x++ {
			// If src is inside the intersection, we blend both
			// pixels using the provided BlendFunc (mode).
			if p := image.Pt(x, y); p.In(inter) {
				img.Set(x, y, mode(dst.At(x, y), src.At(x, y)))
			} else {
				// Else we copy dst pixel to the resulting image.
				img.Set(x, y, dst.At(x, y))
			}
		}
	}
	return img
}

func blendPerChannel(dst, src color.Color, bf func(float64, float64) float64) color.Color {
	d, s := color2rgbaf64(dst), color2rgbaf64(src)
	return rgbaf64{bf(d.r, s.r), bf(d.g, s.g), bf(d.b, s.b), d.a}
}

// Blending modes supported by Photoshop in order.
/*-------------------------------------------------------*/

// Darken ...
func Darken(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, darken)
}
func darken(d, s float64) float64 {
	return math.Min(d, s)
}

// Multiply ...
func Multiply(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, multiply)
}
func multiply(d, s float64) float64 {
	return s * d / max
}

// ColorBurn ...
func ColorBurn(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, colorBurn)
}
func colorBurn(d, s float64) float64 {
	if s == 0.0 {
		return s
	}
	return math.Max(0.0, max-((max-d)*max/s))
}

// LinearBurn ...
func LinearBurn(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, linearBurn)
}
func linearBurn(d, s float64) float64 {
	if (s + d) < max {
		return 0.0
	}
	return s + d - max
}

// DarkerColor ...
func DarkerColor(dst, src color.Color) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	if s.r+s.g+s.b > d.r+d.g+d.b {
		return dst
	}
	return src
}

// Lighten ...
func Lighten(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, lighten)
}
func lighten(d, s float64) float64 {
	return math.Max(d, s)
}

// Screen ...
func Screen(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, screen)
}
func screen(d, s float64) float64 {
	return s + d - s*d/max
}

// ColorDodge ...
func ColorDodge(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, colorDodge)
}
func colorDodge(d, s float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * max / (max - s)))
}

// LinearDodge ...
func LinearDodge(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, linearDodge)
}
func linearDodge(d, s float64) float64 {
	return math.Min(s+d, max)
}

// LighterColor ...
func LighterColor(dst, src color.Color) color.Color {
	s, d := color2rgbaf64(src), color2rgbaf64(dst)
	if s.r+s.g+s.b > d.r+d.g+d.b {
		return src
	}
	return dst
}

/*-------------------------------------------------------*/

// Overlay ...
func Overlay(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, overlay)
}
func overlay(d, s float64) float64 {
	if d < mid {
		return 2 * s * d / max
	}
	return max - 2*(max-s)*(max-d)/max
}

// SoftLight ...
func SoftLight(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, softLight)
}
func softLight(d, s float64) float64 {
	return (d / max) * (d + (2*s/max)*(max-d))
}

// HardLight ...
func HardLight(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, hardLight)
}
func hardLight(d, s float64) float64 {
	if s > mid {
		return d + (max-d)*((s-mid)/mid)
	}
	return d * s / mid
}

// VividLight ...
func VividLight(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, vividLight)
}
func vividLight(d, s float64) float64 {
	if s < mid {
		return colorBurn((2 * s), d)
	}
	return colorDodge((2 * (s - mid)), d)
}

// LinearLight ...
func LinearLight(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, linearLight)
}
func linearLight(d, s float64) float64 {
	if s < mid {
		return linearBurn((2 * s), d)
	}
	return linearDodge((2 * (s - mid)), d)
}

// PinLight ...
func PinLight(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, pinLight)
}
func pinLight(d, s float64) float64 {
	if s < mid {
		return darken((2 * s), d)
	}
	return lighten((2 * (s - mid)), d)
}

// HardMix ...
func HardMix(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, hardMix)
}
func hardMix(d, s float64) float64 {
	if vividLight(d, s) < mid {
		return 0.0
	}
	return max
}

/*-------------------------------------------------------*/

// Difference ...
func Difference(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, difference)
}
func difference(d, s float64) float64 {
	return math.Abs(s - d)
}

// Exclusion ...
func Exclusion(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, exclusion)
}
func exclusion(d, s float64) float64 {
	return s + d - s*d/mid
}

// Substract ...
func Substract(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, substract)
}
func substract(d, s float64) float64 {
	if d-s < 0.0 {
		return 0.0
	}
	return d - s
}

// Divide ...
func Divide(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, divide)
}
func divide(d, s float64) float64 {
	return (d*max)/s + 1.0
}

// Blending modes that use HSL color model transformations.
/*-------------------------------------------------------*/

// Hue ...
func Hue(dst, src color.Color) color.Color {
	s := rgb2hsl(src)
	if s.s == 0.0 {
		return dst
	}
	d := rgb2hsl(dst)
	return hsl2rgb(s.h, d.s, d.l)
}

// Saturation ...
func Saturation(dst, src color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(d.h, s.s, d.l)
}

// Color ...
func Color(dst, src color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(s.h, s.s, d.l)
}

// Luminosity ...
func Luminosity(dst, src color.Color) color.Color {
	s := rgb2hsl(src)
	d := rgb2hsl(dst)
	return hsl2rgb(d.h, d.s, s.l)
}

// This blending modes are not implemented in Photoshop
// or GIMP at the moment, but produced their desired results.
/*-------------------------------------------------------*/

// Add ...
func Add(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, add)
}
func add(d, s float64) float64 {
	if s+d > max {
		return max
	}
	return s + d
}

// Reflex (a.k.a. Glow)
func Reflex(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, reflex)
}
func reflex(d, s float64) float64 {
	if s == max {
		return s
	}
	return math.Min(max, (d * d / (max - s)))
}

// Phoenix ...
func Phoenix(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, phoenix)
}
func phoenix(d, s float64) float64 {
	return math.Min(d, s) - math.Max(d, s) + max
}
