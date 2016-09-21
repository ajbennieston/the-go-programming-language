package unitconv

// Angle conversions

import (
    "fmt"
    "math"
)

func (d Degrees) String() string {
    return fmt.Sprintf("%g˚", d)
}

func (r Radians) String() string {
    return fmt.Sprintf("%g rad", r)
}

func DegreesToRadians(d Degrees) Radians {
    return Radians(d / 180.0 * math.Pi)
}

func RadiansToDegrees(r Radians) Degrees {
    return Degrees(r * 180.0 / math.Pi)
}
