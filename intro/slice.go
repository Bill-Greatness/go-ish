package main

import (
	"fmt"
	"strings"
)

type Test struct {
	i int
	b bool
}

func main() {
	// slice is an array without a the length
	q := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(q)

	r := []bool{true, false, false, true, true, false}
	fmt.Println(r)

	s := []Test{
		{2, true},
		{3, false},
		{4, false},
		{8, true},
		{0, false},
	}
	fmt.Println(s)
	// slice defaults
	s = s[:2]
	fmt.Println("Sliced S", s)

	// slicing to give it a length of 0
	s = s[:0]
	fmt.Println(s)

	// nil slices
	var testMe []int
	fmt.Println(testMe, len(testMe), cap(testMe))

	if testMe == nil {
		fmt.Println("nil")
	}

	// using make
	a := make([]int, 5) // make an array with a length of 5
	printSlice("Slice A", a)

	// making slice with length of 0 and a cap (capacity) of 5
	// first argument is the type, second is length, third is capacity.
	b := make([]int, 0, 5)
	printSlice("Slice B", b)

	c := b[:2]
	printSlice("Slice C", c)

	d := c[2:5]
	printSlice("Slice D", d)

	makeTic()

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func makeTic() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// turing of players.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}

	appendToSlice()
}

func appendToSlice() {
	var s []int
	printShort(s)

	// append works on nil slices
	s = append(s, 0)
	printShort(s)

	// append grows as needed.
	s = append(s, 1)
	printShort(s)

	// append more than one element at one time.
	s = append(s, 1, 2, 3, 4)
	printShort(s)
}

func printShort(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
