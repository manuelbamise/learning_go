package main

import (
	"fmt"
)

func main() {
	// x := 10
	// if x > 5 {
	// 	fmt.Println(x)
	// 	x := 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)
	//
	// if n := rand.Intn(10); n==0{
	// 	fmt.Println("This number is too low")
	// }else if n > 5 {
	// 	fmt.Println("This number is too large: ", n)
	// } else {
	// 	fmt.Println("That's a good number: ", n)
	// }
	//
	// complete C-style for
	// for i := 0;i<10;i++{
	// 	fmt.Println(i)
	// }
	// condition-only for statement
	// i := 1
	// for i < 10{
	// 	fmt.Println(i)
	// 	i = i * 2
	// }
	//
	// infinite for loop
	// for {
	// 	fmt.Println("Hello")
	// 	}
	// newSlice := []int{10,28,22,18,50,39}
	//
	// fmt.Println(newSlice)
	//
	// for _,v := range newSlice {
	// 	fmt.Println(v)
	// // }
	// samples := []string{"Hello", "apple_n!"}
	//
	// outer:
	// 	for _, sample := range samples {
	// 		for i, v := range sample {
	// 			fmt.Println(i, v, string(v))
	// 			if v == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 	}

	fmt.Println("Hello, world!")

}

// NOTES
// Each place where a variable declaration occurs is called a block
// A shadowing variable is a variable that has the same name as a variable in a containing block
// Just tips to be careful when naming things the same as identifiers in the universal block
//
// Parentheses aren't put around the condition in Golang's if statement
//
// FOR is the only looping keyword in Golang. and it is used in four formats.
// complete C-style for
// condition-only for
// infinite for
// for-range
//
// GO has a break keyword to break out of an infinite loop
// for-ranged values are copies of the original.., modifying the variable in the loop will not change the value of the compound type
// there are labels in go used to change the control flow of code
//
// Most of the time, the for-range format will be used
// the bset time to use the complete C-style for loop is when iteration is not starting from the first to last element in the compound type.
//
//
