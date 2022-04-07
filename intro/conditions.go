package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// shorted if
func pow(x, n, lim float64) float64 {
	// there is an inline variable before its usage with if
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func getDaySwitch() {
	fmt.Println("When's Saturday ?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today....")
	case today + 1:
		fmt.Println("Tomorrow..")
	default:
		fmt.Println("Too far away")
	}
}
func getSwitchOn() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux..")
	default:
		fmt.Printf("%s\n", os)

	}
}

func getGreeting() {
	t := time.Now()
	//switch withour a consition is the same as switch true
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")

	}
}

func stackDefers (){
	// defer function calls are pused onto a stack with a LIFO order
	fmt.Println("counting")
	for i := 0; i < 10; i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
}


func main() {
	defer fmt.Println("This should be the last code to be executed!")
	stackDefers()

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 4))
	getSwitchOn()
	getDaySwitch()
	getGreeting()
}
