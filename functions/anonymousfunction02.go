package main


import "fmt"

func operation(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	result := operation(10, 20, func(x, y int) int {
		return x - y
	})
	fmt.Println("Result:", result)
}