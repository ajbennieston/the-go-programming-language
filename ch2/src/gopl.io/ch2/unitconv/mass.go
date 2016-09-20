package unitconv

// Mass conversions

import "fmt"

func (kg Kilograms) String() string {
    return fmt.Sprintf("%g kg", kg)
}

func (lb Pounds) String() string {
    return fmt.Sprintf("%g lb", lb)
}

func KilogramsToPounds(kg Kilograms) Pounds {
    return Pounds(kg * 2.20462)
}

func PoundsToKilograms(lb Pounds) Kilograms {
    return Kilograms(lb * 0.453592)
}
