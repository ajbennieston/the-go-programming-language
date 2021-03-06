// findlinks: Prints the links in an HTML document read from standard input.

package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
        os.Exit(1)
    }

    for _, link := range visit(nil, doc) {
        fmt.Println(link)
    }
}

func visit(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                links = append(links, attr.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }
    return links
}
