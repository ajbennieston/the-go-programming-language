// basename: Remove directory components and a .suffix

package main

import (
    "fmt"
    "os"
)

func main() {
    for _, arg := range os.Args[1:] {
        fmt.Printf("%s\n", basename(arg))
    }
}

func basename(s string) string {
    // Discard last '/' and everything before:
    for i := len(s) - 1 ; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }

    // Preserve everything before last '.':
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }
    return s
}
