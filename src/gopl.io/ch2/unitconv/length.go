package unitconv

// Length conversions

import "fmt"

func (f Feet) String() string {
    return fmt.Sprintf("%g'", f)
}

func (m Metres) String() string {
    return fmt.Sprintf("%gm", m)
}

func FeetToMetres(f Feet) Metres {
    return Metres(f * 0.3048)
}

func MetresToFeet(m Metres) Feet {
    return Feet(m * 3.28084)
}

