package utils

import "fmt"

var a int = 12

func VariableDeclaration() {
	 // Explicit declaration
	 fmt.Println(a)
	 
	 var name string
	 name = "Dinsha"
	 fmt.Println("Name:", name)
 
	 // Type inference
	 var age = 23
	 fmt.Println("Age:", age)
 
	 // Short declaration
	 isActive := true
	 fmt.Println("Is Active:", isActive)
 
	 // Multiple variables
	 a, b, c := 1, 2.5, "hello"
	 fmt.Println(a, b, c)
 
	 // Constants
	 const Pi = 3.14
	 fmt.Println("Pi:", Pi)

	

}