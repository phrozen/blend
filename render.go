// Copyright (c) 2012 Guillermo Estrada. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package blend

import (
  "image/draw"
)

// Work in progress.
type Renderer interface {
  draw.Image
  Render(x, y int)
}

// Work in progress.
type ThreadedRenderer interface {
  draw.Image
  ThreadedRender(x, y chan int, done chan bool)
}
