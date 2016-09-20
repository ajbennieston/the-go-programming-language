package unitconv

// Currency conversions - roughly correct at 2016-09-20 21:35 GMT+0100

import "fmt"

func (gbp GBP) String() string {
    return fmt.Sprintf("£%g", gbp)
}

func (eur EUR) String() string {
    return fmt.Sprintf("€%g", eur)
}

func GBPToEUR(gbp GBP) EUR {
    return EUR(gbp * 1.17)
}

func EURToGBP(eur EUR) GBP {
    return GBP(eur * 0.86)
}
