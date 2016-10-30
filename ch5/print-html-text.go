// print-html-text: Print the text content of an HTML document.

package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "print-html-text: %v\n", err)
        os.Exit(1)
    }

    visit(doc)
    fmt.Println()
}

func visit(n *html.Node) {
    if n.Type == html.TextNode {
        switch n.Parent.Data {
        case "script":
            break
        case "style":
            break
        default:
            fmt.Printf("%s ", n.Data)
        }
    }

    if c := n.FirstChild; c != nil {
        visit(c)
    }

    if c := n.NextSibling; c != nil {
        visit(c)
    }
}
