package utils

import (
	"fmt"
	"sort"
	"strings"
)

func Libraries(){

	//strings
	 greetings := "hello friends"

	 fmt.Println(strings.Contains(greetings, "hello")) //it return the boolean output--> true/false

	 fmt.Println(strings.ReplaceAll(greetings,"hello","hi"))
	 fmt.Println(strings.ToUpper(greetings))
	 fmt.Println(strings.Index(greetings,"ll"))
	 fmt.Println(strings.Split(greetings,""))

	 //original string
	 fmt.Println("original sring:",greetings)

	//sort

	number := []int{40,20,10,34,26}
	sort.Ints(number)
	fmt.Println(number)

	index := sort.SearchInts(number, 40)
	fmt.Println(index)

	name := []string{"dinsha", "anu", "vinu", "sree"}
	sort.Strings(name)
	fmt.Println(name)

	fmt.Println(sort.SearchStrings(name,"sree"))
}