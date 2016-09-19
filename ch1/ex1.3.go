// Echo3 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func repeated_concat() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func string_join() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
    start := time.Now()
    repeated_concat()
    t1 := time.Since(start).Seconds()

    start = time.Now()
    string_join()
    t2 := time.Since(start).Seconds()

    fmt.Println(t1)
    fmt.Println(t2)
}
