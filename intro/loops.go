package main

import "fmt" 

func main(){
	sum := 0
	pile := 0
	counter := 1
	for i := 0; i < 20; i++ {
		sum += i 
	}
	fmt.Println(sum)

	// int and post statements are optional... woaaw 
	for ; sum < 20; {
		pile += sum
	}
	fmt.Println(pile)

	// for is go's while 
	for counter < 20{
		counter += counter
	}
	fmt.Println(counter)

	// "for" ever syntax for {}

}