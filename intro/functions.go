package main

import  (
	"fmt"
)

func addNumbers(x, y int) int {
	// parameters could also be defined as x int, y int
	return x + y
}

func swapText (x, y string) (string, string){
	return y , x
}

// using variables 
var i , j int = 1, 2

func split (sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	//naked return
	return 
}

func playFunc() {
	var test1, test2, test3 = true, false, "no!"
	number := addNumbers(5, 9)
	first, second := swapText("Bill", "Greatness")
	fmt.Println("Addition of Numbers is:", number)
	fmt.Println("Swapped Names is: ", first, second)
	// using naked return 
	fmt.Println(split(20))
	// printing from variables. 
	fmt.Println(i, j, test1, test2, test3)
}
