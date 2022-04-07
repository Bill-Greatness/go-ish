package main
import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
type Vertex1 struct {
	X, Y float64
}

func (v Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y * v.Y)
}

func Scale (v *Vertex1, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex1{3, 4}
	fmt.Printf("Before scaling: %+v, Abs:%v\n", v, v.Abs())

	Scale(&v, 5)
	fmt.Printf("After scaling: %+v, Abs:%v\n", v, v.Abs())

	a = f 
	a = &v 

	// in the next line, a doesn't implement Abser 
	a = v 
	fmt.Println(a.Abs())

}