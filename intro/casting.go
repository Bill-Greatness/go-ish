package main

import (
	"fmt"
	"math"
)

//using consts
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))

	//casting digits
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// type inference
	v := 5.6
	//using const: you cannot use the shorthand assignment := for const
	const theConstant = "Bill Greatness"
	fmt.Printf("v is of type: %T\n", v)
	fmt.Printf("%q is the constant\n", theConstant)

	// print funcs from constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
