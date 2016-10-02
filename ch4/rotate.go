// rotate: Rotate a slice.

package main

import "fmt"

func reverse (s []int) {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }
}

// rotate s left by n places:
func rotate(s []int, n int) {
    reverse(s[:n])
    reverse(s[n:])
    reverse(s)
}

func main() {
    a := [...]int{1, 2, 3, 4, 5}
    fmt.Println(a)
    rotate(a[:], 2)
    fmt.Println(a)
    rotate(a[:], 2)
    fmt.Println(a)
}
