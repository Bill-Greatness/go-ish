package main 

import "fmt"

type Person struct{
	Name string 
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main()  {
	a := Person{"James Smith", 42}
	z := Person{"Zoe Pebble", 200}
	fmt.Println(a, z)
}