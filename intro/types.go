package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func zeroValues () {
	var i int 
	var f float64 
	var b bool 
	var s string 
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func printTypes () {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// << operator 1 << 4 reads 1 x 2 (4 times)
	// >> operator 1 >> 4 reads 1 divided by 2 (4 times)
	fmt.Println(1 << 2)
}
