// test for 'popcount' package
package main

import (
    "fmt"
    "os"
    "strconv"
    "gopl.io/ch2/popcount"
)

func main() {
    for _, arg := range os.Args[1:] {
        val, err := strconv.ParseUint(arg, 0, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "test-popcount: Error converting argument: %s\n", arg)
            continue
        }

        fmt.Printf("The population count of %v is %d\n", val, popcount.PopCount(val))
    }
}
