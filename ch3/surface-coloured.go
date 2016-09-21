// surface-coloured: Compute an SVG rendering of a 3-D surface function.

package main

import (
    "fmt"
    "io"
    "math"
    "os"
)

func main() {
    width := 600
    height := 320
    cells := 100
    xyrange := 30.0
    angle := math.Pi / 6.0

    canvas := makeCanvas(width, height, cells, xyrange, angle)
    surfacePlot(canvas, os.Stdout)
}

// Canvas properties
type Canvas struct {
    width    int
    height   int
    cells    int
    xyrange  float64
    angle    float64

    xyscale  float64
    zscale   float64
    cosAngle float64
    sinAngle float64
}

func makeCanvas(width int, height int, cells int, xyrange float64, angle float64) Canvas {
    xyscale := float64(width) / 2 / xyrange
    zscale := float64(height) * 0.4
    cosAngle := math.Cos(angle)
    sinAngle := math.Sin(angle)
    return Canvas{width, height, cells, xyrange, angle, xyscale, zscale, cosAngle, sinAngle}
}

var palette = [...]string { // blue .. red
    "#0019bf",
    "#1616ac",
    "#2d1499",
    "#441186",
    "#5b0f73",
    "#720C81",
    "#8904ae",
    "#a0073b",
    "#b70528",
    "#ce0215",
    "#e50003"}

func colorMap(z, min, max float64) string {
    // Map one numeric range onto another:
    idx := (float64(len(palette))/(max - min)) * (z - min)
    // Truncate and use as index in palette:
    return palette[int(idx)]
}

// A single point. (x, y) are the canvas coordinates, z is the function height.
type Point struct {
    x float64
    y float64
    z float64
}

// A polygon defined by four points.
type Polygon struct {
    a Point
    b Point
    c Point
    d Point
}

func surfacePlot(canvas Canvas, out io.Writer) {
    // Populate canvas data
    fmt.Fprintf(out,
               "<svg xmlns='http://www.w3.org/2000/svg' " +
               "style='stroke: grey; fill: white; stroke-width: 0.7' " +
               "width='%d' height='%d'>", canvas.width, canvas.height)

    var polygons []Polygon
    zmin := 0.0
    zmax := 0.0

    for i := 0; i < canvas.cells; i++ {
        for j := 0; j < canvas.cells; j++ {
            a, err := corner(i + 1, j, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, a.z)
            zmax = math.Max(zmax, a.z)
            b, err := corner(i, j, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, b.z)
            zmax = math.Max(zmax, b.z)
            c, err := corner(i, j + 1, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, c.z)
            zmax = math.Max(zmax, c.z)
            d, err := corner(i + 1, j + 1, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, d.z)
            zmax = math.Max(zmax, d.z)
            polygons = append(polygons, Polygon{a, b, c, d})
        }
    }

    for _, p := range polygons {
        meanHeight := (p.a.z + p.b.z + p.c.z + p.d.z) / 4.0
        color := colorMap(meanHeight, zmin, zmax)
        fmt.Fprintf(out,
                   "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
                   p.a.x, p.a.y,
                   p.b.x, p.b.y,
                   p.c.x, p.c.y,
                   p.d.x, p.d.y,
                   color)
    }

    fmt.Fprintf(out, "</svg>\n")
}

func corner(i, j int, canvas Canvas) (Point, bool) {
    // Find point (x, y) at corner of cell (i, j)
    x := canvas.xyrange * (float64(i) / float64(canvas.cells) - 0.5)
    y := canvas.xyrange * (float64(j) / float64(canvas.cells) - 0.5)

    // Compute surface height z
    z := f(x, y)
    if math.IsInf(z, 0) || math.IsNaN(z) {
        return Point{}, true
    }

    // Project (x, y, z) isometrically onto 2D SVG canvas (sx, sy)
    sx := float64(canvas.width) / 2 + (x - y) * canvas.cosAngle * canvas.xyscale
    sy := float64(canvas.height) / 2 + (x + y) * canvas.sinAngle * canvas.xyscale - z * canvas.zscale
    return Point{sx, sy, z}, false
}

func f(x,y float64) float64 {
    r := math.Hypot(x, y) // distance from (0, 0)
    return math.Sin(r) / r
}

// Eggbox:
func f2(x, y float64) float64 {
    return (math.Sin(x) + math.Cos(y)) / 10.
}
