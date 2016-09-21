package unitconv

// Energy conversions

import "fmt"

func (j Joules) String() string {
    return fmt.Sprintf("%g J", j)
}

func (kcal Kcal) String() string {
    return fmt.Sprintf("%g Kcal", kcal)
}

func JoulesToKcal(j Joules) Kcal {
    return Kcal(j * 0.000239006)
}

func KcalToJoules(kcal Kcal) Joules {
    return Joules(kcal * 4184)
}
