// convert: string to numeric conversions with default values on failure.
package convert

import "strconv"

// Convert a string to an integer. If the conversion fails, return fallback.
func ConvertInt(s string, fallback int) int {
    if val, err := strconv.Atoi(s); err != nil {
        return fallback
    } else {
        return val
    }
}

// Convert a string to a float64. If the conversion fails, return fallback.
func ConvertFloat64(s string, fallback float64) float64 {
    if val, err := strconv.ParseFloat(s, 64); err != nil {
        return fallback
    } else {
        return val
    }
}

