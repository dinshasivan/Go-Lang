package utils

import "fmt"

func Conditions(){

	age := 45

	fmt.Println(age <= 45)
	fmt.Println(age >= 45)
	fmt.Println(age == 45)
	fmt.Println(age != 45)

	if age < 30 {
		fmt.Println(" is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	//break and continue
	names := []string{"anu","sree","dinsha","vinu"}

	for index, value := range names {
		if index == 1{
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 2{
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n",index, value)
	}

}