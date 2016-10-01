// stack: demonstrate use of a slice as a stack.

package main

import "fmt"

func push(stack *[]int, val int) {
    *stack = append(*stack, val)
}

func pop(stack *[]int){
    *stack = (*stack)[:len(*stack) - 1]
}

func top(stack []int) int {
    return stack[len(stack) - 1]
}

func main() {
    stack := make([]int, 0, 0)
    push(&stack, 1)
    fmt.Printf("%v\n", stack)
    push(&stack, 2)
    fmt.Printf("%v\n", stack)
    pop(&stack)
    fmt.Printf("%v\n", stack)
    fmt.Printf("%v\n", top(stack))
    pop(&stack)
    fmt.Printf("%v\n", stack)
}
