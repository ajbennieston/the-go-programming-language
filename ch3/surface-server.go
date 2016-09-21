// surface-server: SVG rendering of surface plots.
package main

import (
    "log"
    "math"
    "net/http"
    "strconv"
    "gopl.io/ch3/surfaceplot"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    width := 600
    height := 320
    cells := 100
    xyrange := 30.0
    angle := math.Pi / 6.0

    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        if len(v) == 0 {
            continue
        }

        switch k {
        case "width":
            width = convertInt(v[0], width)
        case "height":
            height = convertInt(v[0], height)
        case "cells":
            cells = convertInt(v[0], cells)
        case "xyrange":
            xyrange = convertFloat64(v[0], xyrange)
        case "angle":
            angle = convertFloat64(v[0], angle)
        }
    }

    w.Header().Set("Content-Type", "image/svg+xml")
    canvas := surfaceplot.MakeCanvas(width, height, cells, xyrange, angle)
    surfaceplot.SurfacePlot(f, canvas, w)
}

func convertInt(s string, fallback int) int {
    if val, err := strconv.Atoi(s); err != nil {
        return fallback
    } else {
        return val
    }
}

func convertFloat64(s string, fallback float64) float64 {
    if val, err := strconv.ParseFloat(s, 64); err != nil {
        return fallback
    } else {
        return val
    }
}

func f(x,y float64) float64 {
    r := math.Hypot(x, y) // distance from (0, 0)
    return math.Sin(r) / r
}

