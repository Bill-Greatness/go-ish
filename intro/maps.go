package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	// you can just omit the type names infront of the keys. "B":{ere, 343}
	"Bell Labs": Vertex{
		-33.464424, 83.00045,
	},
	"Google": Vertex{
		32.0034, 30.3443,
	},
}

func main() {
	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	-33.464424, 83.00045,
	// }
	fmt.Println(m)
	answerMap()
}

func answerMap() {
	m := make(map[string]int)
	m["Answer"] = 33
	fmt.Println("Value Now:", m["Answer"])

	m["Answer"] = 41
	fmt.Println("Value Now:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("Deleted Value Now:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The Value:", v, "Present?", ok)

}
