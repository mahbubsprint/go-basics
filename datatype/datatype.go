package main

import "fmt"

func main() {
	var Y int16 = 32767
	var Z int16 = 32768 // overflow
    fmt.Println(Y+2, Y-2,  Z)
}