package utils

import (
	"fmt"
	"strings"
)

func MultipleReturnValue(n string) (string, string){
	s := strings.ToUpper(n)
	names := strings.Split(s,"")

	var initials [] string
	for _, v := range names{
		initials = append(initials, v[:1])
	}

	if len(initials)>1 {
		return initials[0], initials[1]
	} 
	return initials[0], "_"
}

func Sum(numbers ...int){
	total := 0
	for _, num := range numbers{
		total += num
	}
	fmt.Println("Sum", total)
}