package main

import (
    "math"
    "unsafe"
    "github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

func putpixel(surface *sdl.Surface, x, y uint16, pixel uint32) () {
    bpp := surface.Format.BytesPerPixel
    p := uintptr(surface.Pixels) + uintptr(y*surface.Pitch + x*uint16(bpp))
    ptr := unsafe.Pointer(p)
    switch bpp {
    case 1:
        *((*uint32)(ptr)) = pixel
    case 2:
        *((*uint32)(ptr)) = pixel
    case 3:
        *((*uint32)(ptr)) = pixel
    case 4:
        *((*uint32)(ptr)) = pixel
    }
}

func DrawLine(surface *sdl.Surface, x0, y0, x1, y1 int) {
    dx := math.Abs(float64(x1-x0))
    dy := math.Abs(float64(y1-y0))

    var sx, sy int

    if x0 < x1 {
        sx = 1
    } else {
        sx = -1
    }

    if y0 < y1 {
        sy = 1
    } else {
        sy = -1
    }

    err := dx - dy

    for {
        putpixel(surface, uint16(x0), uint16(y0), 0)
        if x0 == x1 && y0 == y1 {
            break
        }

        e2 := 2*err
        if e2 > -dy {
            err = err - dy
            x0 = x0 + sx
        }

        if e2 < dx {
            err = err + dx
            y0 = y0 + sy
        }
    }
}
