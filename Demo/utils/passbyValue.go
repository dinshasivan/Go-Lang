package utils

import "fmt"

func passByValue(x int){ 

	//pass by value
	x = x * 2
	fmt.Println("inside function:", x) //This changes only the local copy



}
type Person struct{
	name string
	age int
}

func changeName(p Person){ // pass by value in truct
	p.name = "Arya"
	fmt.Println("inside function:",p.name)
}

func changeValue(x *int){
	*x = *x * 2
}

func PassByValue(){
	num := 10
	fmt.Println("before function:",num)

	passByValue(num)
	fmt.Println("after function call:",num)

	
	person := Person{name: "Dinsha", age: 23}
	fmt.Println("before function call:",person.name)

	changeName(person)
	fmt.Println("after function call:",person.name)

	value := 20
	fmt.Println("before function:",num)

	changeValue(&value)// pass the address of variable
	fmt.Println("after function call:",num)
}