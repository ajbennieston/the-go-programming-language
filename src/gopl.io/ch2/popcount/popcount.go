// popcount: population count for a 64-bit integer

package popcount

// pc[i] is the population count of i

var pc[256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}

// PopCountLoop uses a loop to compute the population count.
func PopCountLoop(x uint64) int {
    var accumulator int
    for i := uint(0); i < 8; i++ {
        accumulator += int(pc[byte(x>>(i*8))])
    }
    return accumulator
}

// PopCountShift uses right-shift to compute the population count.
func PopCountShift(x uint64) int {
    var accumulator int
    for i := 0; i < 64; i++ {
        if (x&1) == 1 {
            accumulator++
        }
        x >>= 1
    }
    return accumulator
}

// PopCountRightmostNonzero counts the number of non-zero bits
// using a clever trick.
func PopCountRightmostNonzero(x uint64) int {
    var accumulator int
    for x != 0 {
        accumulator++
        x &= (x - 1)
    }
    return accumulator
}
