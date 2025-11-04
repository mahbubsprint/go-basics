// Go program to illustrate
// concept of variable
package main
import "fmt"
 
func main() {
 
    // Multiple variables of the same type
    // are declared and initialized
    // in the single line
//     var myvariable11, myvariable22, myvariable33 int = 2, 454, 67
 
//    // Display the values of the variables
//    fmt.Printf("The value of myvariable11 is : %d\n",
//                                        myvariable11)

//    fmt.Printf("The value of myvariable22 is : %d\n",
//                                        myvariable22)

//    fmt.Printf("The value of myvariable33 is : %d",
//                                       myvariable33)

// 	fmt.Println("\nMultiple variables of different types")
// // are declared and initialized in the single line
// var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56

// // Display the value and 
// // type of the variables
// fmt.Printf("The value of myvariable1 is : %d\n",
//                                     myvariable1)

// fmt.Printf("The type of myvariable1 is : %T\n",
//                                    myvariable1)

// fmt.Printf("\nThe value of myvariable2 is : %s\n",
//                                      myvariable2)

// fmt.Printf("The type of myvariable2 is : %T\n",
//                                    myvariable2)

// fmt.Printf("\nThe value of myvariable3 is : %f\n",
//                                       myvariable3)

// fmt.Printf("The type of myvariable3 is : %T\n",
//                                    myvariable3)
// Using short variable declaration
// Multiple variables of same types
// are declared and initialized in 
// the single line
myvar1, myvar2, myvar3 := 800, 34, 56

// Display the value and 
// type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %d\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %d\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}