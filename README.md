
# blend (WIP)

Image processing library and rendering toolkit for Go.

This library is compatible with Go1.

### Installation:

```
go get github.com/Phrozen/blend
```

### Usage:
```
import "github.com/Phrozen/blend"
```

Use this convenience function:
```go
func BlendImage(source, dest image.Image, mode BlendFunc) (image.Image, error) {
  ...
}
```

For example:

```go
//let's say we already read 2 images 'source' and 'destination'
img, err := BlendImage(source, destination, blend.COLOR_BURN)
if err != nil {
  panic(err)
}
// Save img or blend it again.
```

Easily extensible as it uses the standard library interfaces from 'image' and 'image/color'.

```go
type BlendFunc func(float64, float64) float64
```

The library uses _float64_ for precision, math operations, and conversions to the 'image' interfaces. A __BlendFunc__ is applied to each channel (RGB) of an image (Alpha channel is not utilized atm).

Example:

```go
// This will just substract the Red, Green, and Blue channels of 2 images.
// Pretty useful for finding differences between similar images.
// (This is the actual implementation of the blend.DIFFERENCE blending mode)
func Difference(s, d float64) float64 {
  return math.Abs(s - d)
}

img, err := BlendImage(source, destination, Difference)
```

At the moment it supports the following blending modes:

+ Multiply
+ Screen
+ Overlay
+ Soft Light
+ Hard Light
+ Color Dodge
+ Color Burn
+ Linear Dodge
+ Linear Burn
+ Darken
+ Lighten
+ Difference
+ Exclusion
+ Reflex
+ Linear Light
+ Pin Light
+ Vivid Light
+ Hard Mix

Check the examples directory for more on blending modes.

### More features to come.

#License
### Copyright (c) 2012 Guillermo Estrada

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
