package main

import (
	"Demo/utils"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")

	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument to execute a specific function.")
		return
	}

	switch os.Args[1] {
	case "1":
		utils.Helloworld()
	case "2":
		utils.VariableDeclaration()
	case "3":
		utils.FormatingString()
	case "4":
		utils.ArrayOperation()
	case "5":
		utils.Libraries()
	case "6":
		utils.Loops()
	case "7":
		utils.Conditions()
	case "8":
		utils.Functions("dinsha")
	case "9":
		utils.CycleName([]string{"cloud", "tifa", "barret"}, utils.Functions)
	case "10":
		a1 := utils.CircleArea(10.5)
		a2 := utils.CircleArea(15)
		fmt.Printf("Circle 1: %0.3f, Circle 2: %0.3f\n", a1, a2)
	case "11":
		fn, sn := utils.MultipleReturnValue("dinsha sivan")
		fmt.Println(fn, sn)
	case "12":
		utils.Sum(1, 2, 3, 4, 5)
	case "13":
		utils.MapOerations()
	case "14":
		utils.PassByValue()
	case "15":
		utils.ReceiverFunction()
	default:
		fmt.Println("Invalid argument. Please provide a valid function name.")
	}
}
