
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
func Blend(dst, src image.Image, mode BlendFunc) (image.Image, error) {
  ...
}
// src is the top layer, dst is the bottom layer or image to be applied to.
```

For example:

```go
import "github.com/phrozen/blend"

// Read two images 'source' and 'destination'

// Blend source (top layer) into destination (bottom layer)
// using Color Burn blending mode.
img, err := blend.Blend(destination, source, blend.ColorBurn)
if err != nil {
  panic(err)
}

// Save img or blend it again applying another blend mode.
img, err := blend.Blend(img, source, blend.Screen)
```

Can be easily extended as it uses the standard library interfaces from **'image'** and **'image/color'**.

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
