// squash-spaces: Replace adjacent Unicode spaces with a single ASCII space in
// a UTF-8 string.

package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
)

func squash(s []byte) []byte {
    in_idx := 0
    out_idx := 0
    inSpace := false
    for in_idx < len(s) {
        r, size := utf8.DecodeRune(s[in_idx:])
        if unicode.IsSpace(r) {
            if !inSpace {
                s[out_idx] = ' '
                out_idx++
            }
            inSpace = true
        } else {
            inSpace = false
            copy(s[out_idx:], s[in_idx:in_idx + size])
            out_idx += size
        }
        in_idx += size
    }
    return s[:out_idx]
}

func main() {
    s := "Hello,    world! 𠜎  𠱓 ."
    s2 := string(squash([]byte(s)))
    fmt.Println(s2)
}
