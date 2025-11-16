package main

import "fmt"

func anonymousFunction(a, b int) int {
	return a + b;
}
func receivedAnonymousFunction(f func(a, b int) int) {
 fmt.Println("Result from anonymous function:", f(5, 10))
}

func anonymousFunctionReciever(i func(x, y string) string) {
	result := i("Hello", "World")
	fmt.Println("Result from anonymous function with string:", result)
}

// main demonstrates using anonymous function values and passing them to receivers.
// It forwards a predeclared anonymousFunction to receivedAnonymousFunction, then
// declares an inline function that concatenates two strings and passes it to
// anonymousFunctionReciever. The example highlights creating function literals
// and using higher-order functions to accept and invoke those function values.
func main() {
	receivedAnonymousFunction(anonymousFunction)
	anonymousFunc := func(x, y string) string {
		return x + " " + y
	}
	anonymousFunctionReciever(anonymousFunc)
}