//Compute SHA256 hash of some calues.
package main

import (
    "crypto/sha256"
    "encoding/binary"
    "fmt"
)

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%x\n%x\n%t\n%T\n%d bits differ\n",
               c1,
               c2,
               c1 == c2,
               c1,
               countDifferingBits(c1, c2))
}

func countDifferingBits(p1, p2 [32]byte) int {
    // Accumulator for counting bit differences
    count := 0
    // Convert each bit pattern into 4 64-bit unsigned integers
    for i := 0; i < 32; i += 8 {
        i1 := binary.LittleEndian.Uint64(p1[i:i+8])
        i2 := binary.LittleEndian.Uint64(p2[i:i+8])
        // xor leaves 1s for only the differing bits
        delta := i1 ^ i2
        for delta != 0 { // count the number of 1s
            delta = delta & (delta - 1)
            count++
        }
    }
    return count
}
