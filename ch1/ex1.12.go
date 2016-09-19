// server2 is a minimal echo server with request counts.

package main

import (
    "fmt"
    "image"
    "image/color"
    "image/gif"
    "io"
    "log"
    "math"
    "math/rand"
    "net/http"
    "strconv"
    "sync"
)

var mu sync.Mutex
var count int

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()

    // Set some default values:
    cycles := 5   // number of complete x oscillator revolutions
    res := 0.001  // angular resolution
    size := 100   // image canvas covers [-size..+size]
    nframes := 64 // number of animation frames
    delay := 8    // delay between frames in 10 ms units

    // Use values from the request, if available:
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        if len(v) == 0 {
              continue
        }

        switch k {
        case "cycles":
            cycles = convertInt(v[0], cycles)
        case "res":
            res = convertFloat64(v[0], res)
        case "size":
            size = convertInt(v[0], size)
        case "nframes":
            nframes = convertInt(v[0], nframes)
        case "delay":
            delay = convertInt(v[0], delay)
        }
    }

    lissajous(w, cycles, res, size, nframes, delay)
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

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}


var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference

    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size + 1, 2*size + 1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*float64(size) + 0.5),
                              size+int(y*float64(size) + 0.5),
                              blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
