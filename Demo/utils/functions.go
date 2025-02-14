package utils

import (
	"fmt"
	"math"
)

func Functions(n string){
	fmt.Printf("Good morning %v \n",n)
}

func Bye(n string){
	fmt.Printf("Goodbye %v \n",n)
}

func CycleName(n []string, f func(string)){ //passsing arguments as array/slice and function
	for _, v := range n{
		f(v)
	}
}

func CircleArea(r float64) float64{
	return math.Pi * r * r
}