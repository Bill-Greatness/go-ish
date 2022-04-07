package main 

import (
	"fmt"
	"io"
	"strings"
	"image"
)

func main(){
	r := strings.NewReader("Bill Greatness!")

	b := make([]byte, 12)

	for {
		n, err := r.Read(b)
		fmt.Printf("n= %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b:[:n] = %q\n", b[:n])
		if err == io.EOF{
			break;
		}
	}
	getImage()
}

func getImage () {
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}