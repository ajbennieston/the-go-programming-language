// rotate: Rotate a slice in O(n) time complexity.

package main

import "fmt"

// Time complexity O(n)
// Space requirements O(1)
// Method 3 at http://www.geeksforgeeks.org/array-rotation
func rotate(s []int, d int) {
    n := len(s)
    for i := 0; i < gcd(d, n); i++ {
        temp := s[i]
        j := i
        for {
            k := j + d
            if k >= n {
                k = k - n
            }
            if k == i {
                break
            }
            s[j] = s[k]
            j = k
        }
        s[j] = temp
    }
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    } else {
        return gcd(b, a % b)
    }
}

func main() {
    a := [...]int{1, 2, 3, 4, 5}
    fmt.Println(a)
    rotate(a[:], 2)
    fmt.Println(a)
    rotate(a[:], 2)
    fmt.Println(a)
}
