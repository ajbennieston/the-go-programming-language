// test for 'popcount' package
package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
    "gopl.io/ch2/popcount"
)

func main() {
    for _, arg := range os.Args[1:] {
        val, err := strconv.ParseUint(arg, 0, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "test-popcount: Error converting argument: %s\n", arg)
            continue
        }

        calcTime(val, popcount.PopCount)
        calcTime(val, popcount.PopCountLoop)
        calcTime(val, popcount.PopCountShift)
        calcTime(val, popcount.PopCountRightmostNonzero)
        fmt.Printf("--\n")
    }
}

func calcTime(val uint64, f func(uint64)int) float64 {
    start := time.Now()
    // Run it a few times for timing...
    var result int
    for i := 0; i < 1000; i++ {
        result = f(val)
    }
    dt := time.Since(start).Seconds()
    fmt.Printf("[%g] %v has %d non-zero bits.\n", dt, val, result)
    return dt
}
