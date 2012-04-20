# blend

Image processing library and rendering toolkit for Go. (WIP)

## Installation:

This library is compatible with Go1.

```
go get github.com/phrozen/blend
```

## Usage:
```
import "github.com/phrozen/blend"
```

Use this convenience function:

```go
func BlendNewImage(dst, src image.Image, mode BlendFunc) image.Image {
  ...
}
// src is the top layer, dst is the bottom layer.
```

For example:

```go
import "github.com/phrozen/blend"

// Read two images 'source' and 'destination' (image.Image)

// Blend source (top layer) into destination (bottom layer)
// using Color Burn blending mode.
img1 := blend.BlendNewImage(destination, source, blend.ColorBurn)


// Save img or blend it again applying another blend mode.
img2 := blend.BlendNewImage(img1, source, blend.Screen)
```

If you want to apply the Blend Mode to an image and modify it without returning a copy, you must provide a mutable image type, one that implements **'draw.Image'** interface. Use this function.

```go
func BlendImage(dst draw.Image, src image.Image, mode BlendFunc) {
  ...
}
// src is the top layer, dst is the bottom layer and image that will be applied to.
```

This function is faster as it does not copy the contents of the original image and applies the Blend Mode just to the intersection of both layers. Most images returned by the encoders of the standard library are already mutable as they implement the **'draw.Image'** interface, but you will have to apply and interface/type assertion first. 

*(Note: jpeg decoder returns color images in YCbCr color mode that does not implement **'draw.Image**', PNG decoder returns mostly RGBA family types and should work)*

```go
import "github.com/phrozen/blend"

// Read two images 'source' and 'destination' (image.Image)

dst, ok := destination.(draw.Image)
if ok {
  blend.BlendImage(dst, source, blend.ColorBurn)
  blend.BlendImage(dst, source, blend.Screen)
}
```

The package an be easily extended as it uses the standard library interfaces from **'image'**, **'image/draw'** and **'image/color'**.

```go
type BlendFunc func(dst, src color.Color) color.Color
```

A **BlendFunc** is applied to each color (RGBA) of an image (although included blend modes does not use the Alpha channel atm). Just create your own **BlendFunc** to add custom functionality.


The library uses ***float64*** internally for precision, math operations, and conversions to the **'image'** interfaces. 

#### At the moment it supports the following blending modes:

+ Darken
+ Multiply
+ Color Burn
+ Linear Burn
+ Darker Color

----
+ Lighten
+ Screen
+ Color Dodge
+ Linear Dodge
+ Lighter Color

----
+ Overlay
+ Soft Light
+ Hard Light
+ Vivid Light
+ Linear Light
+ Pin Light
+ Hard Mix

----
+ Difference
+ Exclusion
+ Substract
+ Divide

----
+ Hue
+ Saturation
+ Color
+ Luminosity

----
+ Add
+ Reflex
+ Phoenix

----
**Notes:**

+ *Add, Reflex, and Phoenix modes are not in PSD.*
+ *Vivid Light produces different results than PSD, affects Hard Mix* issue #2
+ *Saturation, Color, and Luminosity modes produce different results than PSD, but the results are either identical to The GIMP or pretty similar.* issue #3

Check the examples directory for more on blending modes.

#### More features to come.

## License:
#### Copyright (c) 2012 Guillermo Estrada

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*MIT License*
