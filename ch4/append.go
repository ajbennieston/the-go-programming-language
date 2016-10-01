// append: Demonstrate appending to a slice, and reallocating if necessary.

package main

import (
    "fmt"
)

func appendInt(x []int, y ...int) []int {
    var z []int
    xlen := len(x)
    xcap := cap(x)
    zlen := xlen + len(y)
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
    copy(z[xlen:], y)
    return z
}

func main() {
    var x, y []int
    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%2d cap=%2d\t%v\n", i, cap(y), y)
        x = y
    }
    y = appendInt(x, 98, 99, 100)
    fmt.Printf("%2d cap=%2d\t%v\n", 10, cap(y), y)
}
