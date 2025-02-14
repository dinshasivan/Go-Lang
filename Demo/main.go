package main

import (
	"Demo/utils"
	"fmt"
	// "fmt"
	// "os"
)

func main() {
	// fmt.Println("Hello world")
	// utils.Helloworld()
	// utils.VariableDeclaration()
	// utils.FormatingString()
	// utils.ArrayOperation()
	// utils.Libraries()
	// utils.Loops()
	// utils.Conditions()
	utils.Functions("dinsha")

	utils.CycleName([]string{"cloud","tifa","barret"},utils.Functions)

	a1 := utils.CircleArea(10.5)
	a2 := utils.CircleArea(15)

	fmt.Println(a1,a2)
	fmt.Printf("cercle 1 is %0.3f and circle 2 is %0.3f \n",a1, a2)

	// if len(os.Args) > 1 {
	// 	if os.Args[1] == "gm" {
	// 		utils.Functions("dinsha")
	// 	} else if os.Args[1] == "gn" {
	// 		utils.Bye("sree")
	// 	}
	// } else {
	// 	fmt.Println("Nothing to do")
	// }

	fn, sn := utils.MultipleReturnValue("dinsha sivan")
	fmt.Println(fn,sn)

	utils.Sum(1,2,3,4,5)
}
