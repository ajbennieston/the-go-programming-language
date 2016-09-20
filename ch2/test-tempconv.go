// test for tempconv package

package main

import (
    "fmt"
    "gopl.io/ch2/tempconv"
)

func main() {
    fmt.Printf("Absolute zero is %v, or %v, or %v (%v, %v)\n",
        tempconv.AbsoluteZeroC,
        tempconv.CToF(tempconv.AbsoluteZeroC),
        tempconv.CToK(tempconv.AbsoluteZeroC),
        tempconv.KToF(tempconv.AbsoluteZeroK),
        tempconv.KToC(tempconv.AbsoluteZeroK))
    fmt.Printf("Freezing point is %v, or %v, or %v (%v, %v, %v)\n",
        tempconv.FreezingC,
        tempconv.CToF(tempconv.FreezingC),
        tempconv.CToK(tempconv.FreezingC),
        tempconv.KToF(tempconv.Kelvin(273.15)),
        tempconv.KToC(tempconv.Kelvin(273.15)),
        tempconv.FToK(tempconv.FreezingF))
}
