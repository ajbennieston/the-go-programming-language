// treesort: Use a binary tree to implement an insertion sort.

package main

import "fmt"

type tree struct {
    value       int
    left, right *tree
}

// Sorts values in place.
func Sort(values []int) {
    var root *tree
    for _, v := range values {
        root = add(root, v)
    }
    appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t * tree) []int {
    if t != nil {
        values = appendValues(values, t.left)
        values = append(values, t.value)
        values = appendValues(values, t.right)
    }
    return values
}

func add(t *tree, value int) *tree {
    if t == nil {
        // Equivalent to return &tree{value: value}.
        t = new(tree)
        t.value = value
        return t
    }

    if value < t.value {
        t.left = add(t.left, value)
    } else {
        t.right = add(t.right, value)
    }
    return t
}

func main() {
    a := [...]int{5, 3, 4, 1, 2, 9, 0, 7, 8, 6}
    fmt.Println(a)
    Sort(a[:])
    fmt.Println(a)
}
