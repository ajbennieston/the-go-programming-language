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
    "sync"
    "storm/convert"
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
            cycles = convert.ConvertInt(v[0], cycles)
        case "res":
            res = convert.ConvertFloat64(v[0], res)
        case "size":
            size = convert.ConvertInt(v[0], size)
        case "nframes":
            nframes = convert.ConvertInt(v[0], nframes)
        case "delay":
            delay = convert.ConvertInt(v[0], delay)
        }
    }

    lissajous(w, cycles, res, size, nframes, delay)
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
