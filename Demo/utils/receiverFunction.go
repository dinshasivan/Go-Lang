package utils

import "fmt"

type Persons struct{
	name string
	age int
}

//Receiver function (method)
func (p Persons) greet(){
	fmt.Println("hello, my name is",p.name)
}

func ReceiverFunction(){
	p := Persons{name: "Dinsha", age: 23}
	p.greet()
}