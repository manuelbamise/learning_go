package main

import "fmt"

func main() {
	// Array example
	// var x =  [3]int{21,23,23}

	//Slice example
	// var x = []int{32, 23, 43, 23, 24}
	// var y = []int{34,5,4,86,88,35}
	// var x []int

	//Multidimensional slice example
	// var x = [][]int (slice of slices)

	// x = append(x, y...)
	// fmt.Println(x, len(x), cap(x))

	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))

	// creating a slice using make
	//
	// x := []string{"A","B","C","D"}
	//
	// fmt.Println(x, len(x), cap(x))
	//
	// y := x[:2]
	//
	// fmt.Println(y, len(y), cap(y))
	// y = append(y,"E")
	//
	//
	// fmt.Println("x:",x)
	// fmt.Println("y:",y)
	x := []string{ "a", "b", "c", "d"}

	y := make([]string, 5)
	num := copy(y,x)

	fmt.Println(y,num)
}

// NOTES
//
// Arrrays are rarely used in go unless you know the exact length needed ahead of time
//
//Slices is better used over arrays beacause slices grow as needed(length is not part of the slice type)
//the zero value for a slice is nil
//Each element assigned in a slice is assigned consecutive memory locations, the length of a slice is the number of consecutive memory locations that have been assigned a value
//
//
//len is an in-built function used to ge tthe size of a slice or array"Hello, world!"
//append is a built-in function used to grow slices
//
//Go is a call by value language. Everytime you pass a parameter to a function , GO
//makes a copy of the value passed in.
//
// the copy command is used to copy the values from a source slice to a destination making the copied slice independent of the source
