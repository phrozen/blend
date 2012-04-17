
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
func Substract(s, d float64) float64 {
  return s-d
}

img, err := BlendImage(source, destination, Substract)
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
