// is_anagram: Determine whether two strings are anagrams.

package main

import (
    "fmt"
    "os"
    "sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func main() {
    n := len(os.Args)
    if n != 3 {
        fmt.Fprintf(os.Stderr,
            "Usage: %s s1 s2\nDetermines whether s1 and s2 are anagrams.\n",
            os.Args[0])
        os.Exit(1)
    }

    s1 := []rune(os.Args[1])
    s2 := []rune(os.Args[2])
    sort.Sort(sortRunes(s1))
    sort.Sort(sortRunes(s2))

    if string(s1) == string(s2) {
        fmt.Printf("Anagrams\n")
    } else {
        fmt.Printf("Not anagrams\n")
    }
}
