package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Bill"
	a[1] = "Greatness"
	fmt.Println(a)

	primes := [6] int{1, 2,3,4,5,6} 
	fmt.Println(primes)
	
	//slicing
	var s [] int = primes[1:4]
	fmt.Println(s)
}
