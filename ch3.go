package main

import "fmt"

func main() {
	// Array example
	// var x =  [3]int{21,23,23}
	
	//Slice example
	var x = []int{32,23,43,23,24}
	
	//Multidimensional slice example
	// var x = [][]int (slice of slices)

	fmt.Println(len(x))

}

// NOTES
//
// Arrrays are rarely used in go unless you know the exact length needed ahead of time
//
//Slices is better used over arrays beacause slices grow as needed(length is not part of the slice type)
//
