// append: Demonstrate appending to a slice, and reallocating if necessary.

package main

import (
    "fmt"
)

func appendInt(x []int, y int) []int {
    var z []int
    xlen := len(x)
    xcap := cap(x)
    zlen := xlen + 1
    if zlen <= xcap {
        // There is room to grow. Extend the slice.
        z = x[:zlen]
    } else {
        // There is insufficient space. Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen
        if zcap < 2 * xlen {
            zcap = 2 * xlen
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }
    z[xlen] = y
    return z
}

func main() {
    var x, y []int
    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d cap=%2d\t%v\n", i, cap(y), y)
        x = y
    }
}
