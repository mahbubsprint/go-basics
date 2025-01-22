package main

import "fmt"

// Animal interface represents any type that can make a sound.
type Animal interface {
	Sound() string
}

// Dog is a type that satisfies the Animal interface.
type Dog struct{}

// Sound returns the sound a dog makes.
func (d Dog) Sound() string {
	return "Woof!"
}

// Cat is a type that satisfies the Animal interface.
type Cat struct{}

// Sound returns the sound a cat makes.
func (c Cat) Sound() string {
	return "Meow!"
}

func main() {
	// Create a slice of animals.
	animals := []Animal{Dog{}, Cat{}}

	// Iterate over the animals and make them sound.
	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}