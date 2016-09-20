// test for tempconv package

package main

import (
    "fmt"
    "gopl.io/ch2/tempconv"
)

func main() {
    fmt.Printf("Absolute zero is %v, or %v\n", tempconv.AbsoluteZeroC, tempconv.CToF(tempconv.AbsoluteZeroC))
}
