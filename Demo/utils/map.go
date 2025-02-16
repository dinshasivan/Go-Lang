package utils

import "fmt"

func MapOerations(){

	//string as key 
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k, v := range menu{
		fmt.Println(k, "-", v)
	}

	//ints as key type
	students := map[int]string{
		1: "dinsha",
		2: "sree",
		3: "anu",
	}
	fmt.Println(students)
	fmt.Println(students[2])

	students[3]= "vinu"

	fmt.Println(students)
}