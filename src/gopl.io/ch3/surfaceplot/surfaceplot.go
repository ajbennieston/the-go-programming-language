// surfaceplot: Compute SVG rendering of a 3D surface function.

package surfaceplot

import (
    "fmt"
    "io"
    "math"
)

// Information about the canvas geometry.
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

// Create a Canvas object with the specified properties.
func MakeCanvas(width, height, cells int, xyrange, angle float64) Canvas {
    xyscale := float64(width) / 2 / xyrange
    zscale := float64(height) * 0.4
    cosAngle := math.Cos(angle)
    sinAngle := math.Sin(angle)
    return Canvas{width, height, cells, xyrange, angle, xyscale, zscale, cosAngle, sinAngle}
}

// Render a function to SVG using the geometry contained in the canvas object.
func SurfacePlot(f func (float64, float64) float64, canvas Canvas, out io.Writer) {
    // Write SVG header:
    fmt.Fprintf(out,
               "<svg xmlns='http://www.w3.org/2000/svg' " +
               "style='stroke: grey; fill: white; stroke-width: 0.7' " +
               "width='%d' height='%d'>", canvas.width, canvas.height)

    // Build an array of polygons, and track the bounds of z:
    var polygons []polygon
    zmin, zmax := 0.0, 0.0
    for i := 0; i < canvas.cells; i++ {
        for j := 0; j < canvas.cells; j++ {
            a, err := corner(f, i + 1, j, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, a.z)
            zmax = math.Max(zmax, a.z)
            b, err := corner(f, i, j, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, b.z)
            zmax = math.Max(zmax, b.z)
            c, err := corner(f, i, j + 1, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, c.z)
            zmax = math.Max(zmax, c.z)
            d, err := corner(f, i + 1, j + 1, canvas)
            if err {
                continue
            }
            zmin = math.Min(zmin, d.z)
            zmax = math.Max(zmax, d.z)
            polygons = append(polygons, polygon{a, b, c, d})
        }
    }

    // Write out each polygon, with a fill colour chosen based on height:
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

    // Write out the SVG closing tag:
    fmt.Fprintf(out, "</svg>\n")
}

