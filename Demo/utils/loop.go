package utils

import "fmt"

 func Loops(){

	//for loop as while loop
	x := 0
	for x < 5{
		fmt.Println("value of x is:", x)
		x++
	}

	//basic for loop
	for i := 0; i < 5; i++{
		fmt.Println("value of i is :",i)
	}

	//looping over array
	names := []string{"anu","sree","dinsha","vinu"}

	for i := 0; i <len(names); i++{
		fmt.Println(names[i])
	}

	for index, value := range names{
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names{
		fmt.Printf("the value is %v \n", value)
	}

	fmt.Println(names)
 }