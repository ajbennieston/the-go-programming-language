// comma: Insert commas into non-negative decimal integer strings.

package main

import (
    "bytes"
    "fmt"
    "os"
)

func main() {
    for _, arg := range os.Args[1:] {
        fmt.Printf("%s\n", comma(arg))
    }
}

func comma(s string) string {
    var buf bytes.Buffer
    n := len(s)
    m := n % 3
    if m != 0 {
        buf.WriteString(s[0:m])
        buf.WriteString(",")
    }
    for i := m; i < n; i += 3 {
        buf.WriteString(s[i: i + 3])
        if i < n - 3 {
            buf.WriteString(",")
        }
    }
    return buf.String()
}

