package main 
import "fmt"

var power = []int{1,2, 4, 8, 16, 32, 64, 128}

func main(){
	printRange()
	makeRange()

}

func printRange(){
	for i, v := range power{
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func makeRange(){
	pwd := make([] int, 10) 
	// get indexes
	for  i:= range pwd{
		pwd[i] = 1 << uint(i)
	}

	// omit ther first / second variable with an underscore (_)
	for _, value := range pwd{
		fmt.Printf("%d\n", value)
	}
}