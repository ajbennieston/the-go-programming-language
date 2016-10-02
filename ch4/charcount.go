// charcount: Count occurrences of Unicode code points.

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    counts := make(map[rune]int) // counts of Unicode code points
    categoryCounts := make(map[string]int) // counts of code points by Unicode category
    var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
    invalid := 0

    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune() // rune, nbytes, error
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        counts[r]++
        utflen[n]++
        category := "unknown"
        switch {
        case unicode.IsControl(r):
            category = "control"
        case unicode.IsDigit(r):
            category = "digit"
        case unicode.IsLetter(r):
            category = "letter"
        case unicode.IsMark(r):
            category = "mark"
        case unicode.IsNumber(r):
            category = "number"
        case unicode.IsPunct(r):
            category = "punctuation"
        case unicode.IsSpace(r):
            category = "space"
        case unicode.IsSymbol(r):
            category = "symbol"
        }
        categoryCounts[category]++
    }

    fmt.Printf("rune\tcount\n")
    fmt.Printf("----\t-----\n")
    for c, n := range counts {
        fmt.Printf("%q\t%4d\n", c, n)
    }
    fmt.Printf("\nlen\tcount\n")
    fmt.Printf("---\t-----\n")
    for i, n := range utflen {
        fmt.Printf("%d\t%4d\n", i, n)
    }
    fmt.Printf("\n")
    for category, count := range categoryCounts {
        fmt.Printf("%12s:\t%4d\n", category, count)
    }
    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}
