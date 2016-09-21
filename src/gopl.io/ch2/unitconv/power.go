package unitconv

// Power conversions

import "fmt"

func (w Watts) String() string {
    return fmt.Sprintf("%g W", w)
}

func (hp Horsepower) String() string {
    return fmt.Sprintf("%g hp", hp)
}

func (ps PS) String() string {
    return fmt.Sprintf("%g PS", ps)
}

func WattsToHorsepower(w Watts) Horsepower {
    return Horsepower(w * 0.00134102)
}

func HorsepowerToWatts(hp Horsepower) Watts {
    return Watts(hp * 745.7)
}

func WattsToPS(w Watts) PS {
    return PS(w * 0.00135962)
}

func PSToWatts(ps PS) Watts {
    return Watts(ps * 735.499)
}

func PSToHorsepower(ps PS) Horsepower {
    return WattsToHorsepower(PSToWatts(ps))
}

func HorsepowerToPS(hp Horsepower) PS {
    return WattsToPS(HorsepowerToWatts(hp))
}
