package utils

import (
	"fmt"
	
)

func FormatingString(){

	name := "dinsha"
	//print
	fmt.Print("hello, ")
	fmt.Print("world")

	//println
	fmt.Println("hello dinsha")
	fmt.Println("hello world")
	fmt.Println("my name is",name)

	//printf (formating string)
	fmt.Printf("my name is %v \n",name)
	fmt.Printf("name is type of %T \n",name)

	//sprintf (save formatted string)

	var str = fmt.Sprintf("my name is %v \n",name)
	fmt.Println("the saved string is",str)
}