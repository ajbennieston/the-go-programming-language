// fetch prints the content found at each specified URL.

package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {

        switch {
            case strings.HasPrefix(url, "https://"):
            case strings.HasPrefix(url, "http://"):
            default:
                url = "http://" + url
        }

        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        _, err = io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Fprintf(os.Stderr, "Status code: %s\n", resp.Status)
    }
}
