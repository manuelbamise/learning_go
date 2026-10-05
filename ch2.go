package main

import (
	"fmt"
)

func main() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i)
	fmt.Println(f)
}

// NOTES
//
// ZERO VALUES
// A default zero value is assigned to a variable declared but not assigned a value
//
// LITERAL
// An explicitly specified number, char,or string. Integer literals(1_384), floating point literal(6.03e34,0x12.34p5)
// rune literals('\141\','\U00000061'), string literals("Greetings and Salutations")
//
// NUMERIC TYPES
// Go has 12 numeric types
//
// Integer types: signed and unsigned from one to eight bytes (8,16,32,64)
// zero value for integer type is 0
//
// Special integer types: byte(uint8), int(32-bit cpu = int32, 64-bit cpu = int64)
// uint(0 or positive)
//
// Floating-point types: Go has two floating point types, float32 & float64
// Dividing a non-zero floating-point value by 0 produces a +inf or -inf
//
// Complex types: Go has first-class support for complex numbers
//
// STRINGS AND RUNES
// The zero value for a string is an empty string
// type rune means int32 just like byte means uint8, the type should be used to clarify intent
// Go does not allow truthiness, converting from another type to bool. If you want to do that you must use comparison operators.
//
// when declaring variable at the package level "var" should be used because := is not legal outside of functions
//
