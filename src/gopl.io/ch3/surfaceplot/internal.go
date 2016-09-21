package surfaceplot

// The definitions in this file are internal to the surfaceplot package and are
// not exported.

import "math"

// (x, y) coordinates in the canvas plane, z original function height.
type point struct {
    x float64
    y float64
    z float64
}

// A polygon defined by four points.
type polygon struct {
    a point
    b point
    c point
    d point
}

// Evaluate f(x, y) for the point corresponding to the cell (i, j)
// and project the result onto the canvas whose geometry is contained in
// the canvas object.
// The boolean part of the return value is true on error, false otherwise.
func corner(f func (float64, float64) float64, i, j int, canvas Canvas) (point, bool) {
    // Find point (x, y) at corner of cell (i, j):
    x := canvas.xyrange * (float64(i) / float64(canvas.cells) - 0.5)
    y := canvas.xyrange * (float64(j) / float64(canvas.cells) - 0.5)

    // Compute surface height z:
    z := f(x, y)
    if math.IsInf(z, 0) || math.IsNaN(z) {
        // Error computing f(x, y). Return an empty point and signal an error:
        return point{}, true
    }

    // Project (x, y, z) isometrically onto 2D SVG canvas (sx, sy):
    sx := float64(canvas.width) / 2 + (x - y) * canvas.cosAngle * canvas.xyscale
    sy := float64(canvas.height) / 2 + (x + y) * canvas.sinAngle * canvas.xyscale -
          z * canvas.zscale

    // Construct a Point from the data, and signal that there was no error:
    return point{sx, sy, z}, false
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
