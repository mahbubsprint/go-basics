package main

import "fmt"



func main() {

	func() {
		fmt.Println("This is an anonymous function")
	}()
	
	value := func() {//anonymous function assigned to a variable
		fmt.Println("This is an anonymous function assigned to a variable")
	}
	value()

	func(a string) {//anonymous function with parameter
		fmt.Println(a)
	}("Learning Go Language")

}