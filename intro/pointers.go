package main 

import "fmt"

type Vertex struct {
	X, Y int
}

func main(){
	 i, j := 42, 20
	 p := &i // 42 
	 fmt.Println("Printing P before change", *p)
	 *p = 21 // 21 intends sets i to 21
	 fmt.Println(i)

	 p = &j // p now 20 
	 *p = *p / 20
	 fmt.Println(j)

	 v := Vertex{1, 2}
	 v.X = 5 
	 fmt.Println(v)

	 pr := &v
	// you can access the value of pr without explicitly dereferencing it... pr.X = Hello
	 // dereferencing

	 //reassigning Vertex 
	 v1 := Vertex{X:1} // y is 0
	 v2 := Vertex{} // both values are zero
	fmt.Println(v1, v2)
	 fmt.Println(pr)
}