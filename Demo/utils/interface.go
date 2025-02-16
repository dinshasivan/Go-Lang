package utils

import "fmt"


// Define an interface
type Speaker interface {
	Speak() string
}

// define an interface with multiple methods
type Shape interface{
	Area() float64
	Perimeter() float64
}

//Rectanlge struct
type Rectanlge struct{
	Width, Height float64
}

//implement the shape interface
func (r Rectanlge) Area() float64{
	return r.Width * r.Height
}

func (r Rectanlge) Perimeter() float64{
	return 2 * (r.Width + r.Height)
}

// Struct implementing the interfac
type Cat struct{}

// Methods implementing the Speaker interfac
func (c Cat) Speak() string {
	return "Meow!"
}

func Interface() {
	// Declare variables of interface type
	var animal Speaker

	animal = Cat{}
	fmt.Println(animal.Speak()) 

	var s Shape = Rectanlge{Width: 5, Height: 10}
	fmt.Println("Area:", s.Area())         // Output: 50
	fmt.Println("Perimeter:", s.Perimeter())
}