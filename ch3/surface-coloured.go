// surface-coloured: Compute an SVG rendering of a 3-D surface function.

package main

import (
    "math"
    "os"
    "gopl.io/ch3/surfaceplot"
)

func main() {
    width := 600
    height := 320
    cells := 100
    xyrange := 30.0
    angle := math.Pi / 6.0

    canvas := surfaceplot.MakeCanvas(width, height, cells, xyrange, angle)
    surfaceplot.SurfacePlot(f, canvas, os.Stdout)
}

func f(x,y float64) float64 {
    r := math.Hypot(x, y) // distance from (0, 0)
    return math.Sin(r) / r
}

// Eggbox:
func f2(x, y float64) float64 {
    return (math.Sin(x) + math.Cos(y)) / 10.
}
