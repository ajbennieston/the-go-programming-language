package unitconv

// Temperature conversions

import "fmt"

// Constants
const (
    AbsoluteZeroK Kelvin = 0
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
    FreezingF Fahrenheit = 32
)

func (k Kelvin) String() string {
    return fmt.Sprintf("%gK", k)
}

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g°F", f)
}

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c * 9 / 5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
    return Kelvin(c - AbsoluteZeroC)
}

func KToC(k Kelvin) Celsius {
    return Celsius(k) + AbsoluteZeroC
}

func FToK(f Fahrenheit) Kelvin {
    return CToK(FToC(f))
}

func KToF(k Kelvin) Fahrenheit {
    return CToF(KToC(k))
}
