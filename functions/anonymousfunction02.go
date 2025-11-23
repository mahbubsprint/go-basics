package main


import "fmt"

// Think of operate as a machine that takes two numbers and a "rule" (function) for combining them.

// •  If you give it a "subtract rule," it subtracts.
// •  If you gave it an "add rule," it would add.
// •  If you gave it a "multiply rule," it would multiply.

// The "rule" here is the anonymous function.

func operation(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	result := operation(10, 20, func(x, y int) int {
		return x - y
	})
	fmt.Println("Result:", result)
}