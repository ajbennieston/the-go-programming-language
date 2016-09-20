// ftoc: convert Fahrenheit to Celsius.

package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        f, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Could not convert argument: %s\n", arg)
            continue
        }
        fmt.Printf("%g°F is %g°C\n", f, fToC(f))
    }
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
