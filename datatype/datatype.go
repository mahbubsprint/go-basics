package main

import "fmt"
import "github.com/shopspring/decimal"

func main() {
	var Y int16 = 32767
	//var Z int16 = 32768 // overflow
	var Z int16 = 3275 // overflow
    fmt.Println(Y+2, Y-2,  Z)

	a:= float32(34.56)
	fmt.Printf("The value of a is %f and type is %T\n", a, a)

}