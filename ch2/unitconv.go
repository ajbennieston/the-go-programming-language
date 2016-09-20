// unitconv: Show conversions of the command-line arguments.

package main

import (
    "fmt"
    "os"
    "strconv"
    "gopl.io/ch2/unitconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        val, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "unitconv: Failed to parse argument: %s\n", arg)
            continue
        }
        showConversions(val)
    }
}

func showConversions(val float64) {
    fmt.Printf("-- Temperature --\n")
    f := unitconv.Fahrenheit(val)
    fmt.Printf("%s == %s\n", f, unitconv.FToC(f))
    fmt.Printf("%s == %s\n", f, unitconv.FToK(f))
    c := unitconv.Celsius(val)
    fmt.Printf("%s == %s\n", c, unitconv.CToF(c))
    fmt.Printf("%s == %s\n", c, unitconv.CToK(c))
    k := unitconv.Kelvin(val)
    fmt.Printf("%s == %s\n", k, unitconv.KToF(k))
    fmt.Printf("%s == %s\n", k, unitconv.KToC(k))

    fmt.Printf("-- Length --\n")
    ft := unitconv.Feet(val)
    fmt.Printf("%s == %s\n", ft, unitconv.FeetToMetres(ft))
    m := unitconv.Metres(val)
    fmt.Printf("%s == %s\n", m, unitconv.MetresToFeet(m))

    fmt.Printf("-- Angle --\n")
    deg := unitconv.Degrees(val)
    fmt.Printf("%s == %s\n", deg, unitconv.DegreesToRadians(deg))
    rad := unitconv.Radians(val)
    fmt.Printf("%s == %s\n", rad, unitconv.RadiansToDegrees(rad))

    fmt.Printf("-- Mass --\n")
    kg := unitconv.Kilograms(val)
    fmt.Printf("%s == %s\n", kg, unitconv.KilogramsToPounds(kg))
    lb := unitconv.Pounds(val)
    fmt.Printf("%s == %s\n", lb, unitconv.PoundsToKilograms(lb))

    fmt.Printf("-- Currency --\n")
    gbp := unitconv.GBP(val)
    fmt.Printf("%s == %s\n", gbp, unitconv.GBPToEUR(gbp))
    eur := unitconv.EUR(val)
    fmt.Printf("%s == %s\n", eur, unitconv.EURToGBP(eur))

    fmt.Printf("-- Energy --\n")
    j := unitconv.Joules(val)
    fmt.Printf("%s == %s\n", j, unitconv.JoulesToKcal(j))
    kcal := unitconv.Kcal(val)
    fmt.Printf("%s == %s\n", kcal, unitconv.KcalToJoules(kcal))

    fmt.Printf("-- Power --\n")
    w := unitconv.Watts(val)
    fmt.Printf("%s == %s\n", w, unitconv.WattsToHorsepower(w))
    fmt.Printf("%s == %s\n", w, unitconv.WattsToPS(w))
    hp := unitconv.Horsepower(val)
    fmt.Printf("%s == %s\n", hp, unitconv.HorsepowerToWatts(hp))
    fmt.Printf("%s == %s\n", hp, unitconv.HorsepowerToPS(hp))
    ps := unitconv.PS(val)
    fmt.Printf("%s == %s\n", ps, unitconv.PSToWatts(ps))
    fmt.Printf("%s == %s\n", ps, unitconv.PSToHorsepower(ps))
}
