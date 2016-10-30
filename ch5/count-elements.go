// count-elements: Print counts of html elements.

package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "count-elements: %v\n", err)
        os.Exit(1)
    }

    elementMap := make(map[string]int)
    for tag, count := range visit(elementMap, doc) {
        fmt.Printf("%s: %d\n", tag, count)
    }
}

func visit(elementCount map[string]int, n *html.Node) map[string]int {
    if n.Type == html.ElementNode {
        elementCount[n.Data] += 1
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        elementCount = visit(elementCount, c)
    }
    return elementCount
}
