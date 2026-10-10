package main

import ("fmt"
			// "math/rand"
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
	i := 1
	for i < 100{
		fmt.Println(i)
		i = i +1
	}

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
