// wordcount: Count the number of occurrences of each word in input.

package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

func main() {
    words := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    input.Split(bufio.ScanWords)
    for input.Scan() {
        word := input.Text()
        words[word]++
    }
    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
        os.Exit(1)
    }

    // Determine highest count:
    maxCount := 0
    for _, c := range words {
        if c > maxCount {
            maxCount = c
        }
    }
    // Number of digits in largest count determines field width:
    numDigits := int(math.Log10(float64(maxCount))) + 1

    for word, count := range words {
        fmt.Printf("%[1]*[2]d\t%[3]s\n", numDigits ,count, word)
    }
}
