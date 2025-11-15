package main

import "fmt"

func sum(numbers ...int) int{
	total := 0
	// for i := 0; i < len(numbers); i++ {
	// 	total += numbers[i]
	// }
	// for _, number := range numbers {//index is ignored
	// 	total += number
	// }
	for i, number := range numbers {//index is ignored
		fmt.Printf("Index: %d, Value: %d\n", i, number)
		total += number
	}
	return total
}

func main() {
	result := sum(2, 3, 4, 5)
	fmt.Printf("The sum is: %d", result)
}