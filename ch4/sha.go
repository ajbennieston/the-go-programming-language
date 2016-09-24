//Compute SHA256 hash of some calues.
package main

import (
    "crypto/sha256"
    "crypto/sha512"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    size := flag.Int("size", 256, "hash size [256, 384 or 512]")
    flag.Parse()

    content, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr,
                    "Error: %v\n", err)
        os.Exit(1)
    }

    switch *size {
    case 256:
        fmt.Printf("%x\n", sha256.Sum256(content))
    case 384:
        fmt.Printf("%x\n", sha512.Sum384(content))
    case 512:
        fmt.Printf("%x\n", sha512.Sum512(content))
    }
}

