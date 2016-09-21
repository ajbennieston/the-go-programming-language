// surface-coloured: Compute an SVG rendering of a 3-D surface function.

package main

import (
    "fmt"
    "math"
)

const (
    width, height = 600, 320            // canvas size in pixels
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angle of x and y axes (30˚)
)

var palette = [...]string {
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

func color_map(height, min, max float64) string {
    x := (float64(len(palette))/(max - min)) * (height - min)
    return palette[int(x)]
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type Point struct {
    x float64
    y float64
    z float64
}

type Polygon struct {
    a Point
    b Point
    c Point
    d Point
}

func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
               "style='stroke: grey; fill: white; stroke-width: 0.7' " +
               "width='%d' height='%d'>", width, height)

    var polygons []Polygon
    zmin := 0.0
    zmax := 0.0

    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            a, err := corner(i + 1, j)
            if err {
                continue
            }
            zmin = math.Min(zmin, a.z)
            zmax = math.Max(zmax, a.z)
            b, err := corner(i, j)
            if err {
                continue
            }
            zmin = math.Min(zmin, b.z)
            zmax = math.Max(zmax, b.z)
            c, err := corner(i, j + 1)
            if err {
                continue
            }
            zmin = math.Min(zmin, c.z)
            zmax = math.Max(zmax, c.z)
            d, err := corner(i + 1, j + 1)
            if err {
                continue
            }
            zmin = math.Min(zmin, d.z)
            zmax = math.Max(zmax, d.z)
            polygons = append(polygons, Polygon{a, b, c, d})
        }
    }
    for _, p := range polygons {
        mean_height := (p.a.z + p.b.z + p.c.z + p.d.z) / 4.0
        color := color_map(mean_height, zmin, zmax) 
        fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
                   p.a.x, p.a.y,
                   p.b.x, p.b.y,
                   p.c.x, p.c.y,
                   p.d.x, p.d.y,
                   color)
    }
    fmt.Printf("</svg>\n")
}

func corner(i, j int) (Point, bool) {
    // Find point (x, y) at corner of cell (i, j)
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    // Compute surface height z
    z := f(x, y)
    if math.IsInf(z, 0) || math.IsNaN(z) {
        return Point{}, true
    }

    // Project (x, y, z) isometrically onto 2D SVG canvas (sx, sy)
    sx := width / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
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
