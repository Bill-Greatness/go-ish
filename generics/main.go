package main

import "fmt"

func sumInts(m map[string]int64) int64 {

	var sum int64

	for _, v := range m {
		sum += v
	}

	return sum
}

func sumFloats(m map[string]float64) float64 {
	var sum float64

	for _, v := range m {
		sum += v
	}

	return sum
}

// Creating a Generic function

type Number interface {
	int64 | float64
}

func sumNumbers[K comparable, V Number](m map[K]V) V {

	var sum V
	for _, v := range m {
		sum += v
	}

	return sum
}

func main() {
	// create a map of string keys and integer values.
	intMap := map[string]int64{
		"f": 5,
		"s": 10,
	}

	floatMap := map[string]float64{
		"f": 6.345,
		"s": 0.331,
	}

	//pass maps to thier functions.
	sumInt := sumInts(intMap)

	sumFlt := sumFloats(floatMap)

	fmt.Printf("Sum of Ints, %d\n", sumInt)
	fmt.Printf("Sum of Floats, %2f", sumFlt)

	gen := sumNumbers(intMap)
	fmt.Println(gen)

}
