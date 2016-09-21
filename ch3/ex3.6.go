// mandelbrot: Emit PNG image of the Mandelbrot fractal.

package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
)

func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height          = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        yPlus := (float64(py)+0.5)/height * (ymax-ymin) + ymin
        yMinus := (float64(py)-0.5)/height * (ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            xPlus := (float64(px)+0.5)/width * (xmax - xmin) + xmin
            xMinus := (float64(px)-0.5)/width * (xmax - xmin) + xmin
            zA := complex(xMinus, yMinus)
            zB := complex(xMinus, yPlus)
            zC := complex(xPlus, yMinus)
            zD := complex(xPlus, yPlus)

            // Image point (px, py) represents complex value z.
            avg := (mandelbrot(zA) + mandelbrot(zB) + mandelbrot(zC) + mandelbrot(zD))/4.0
            if avg > 0 {
                avg += 55
            }
            color := color.RGBA{0x00, 0x00, uint8(avg), 0xff}
            img.Set(px, py, color)
        }
    }
    png.Encode(os.Stdout, img) // NOTE: Ignoring errors.
}

func bluePalette() []color.Color {
    var palette []color.Color
    for b := 0; b < 0xff; b ++ {
        palette = append(palette, color.RGBA{0x00, 0x00, uint8(b), 0xff})
    }
    return palette
}

var palette = bluePalette()

func mandelbrot(z complex128) int {
    const iterations = 200
    const contrast = 15
    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v * v + z
        if cmplx.Abs(v) > 2 {
            return int(n)
        }
    }
    return 0
}
