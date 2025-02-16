package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserInput() {
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // read user input
	fmt.Println("Hello", name)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name:")
	names,_ := reader.ReadString('\n')
	names = strings.TrimSpace(names)
	fmt.Println(names)

	fmt.Println("enter your age:")
	input, _:=reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(input))

	if  err != nil {
		fmt.Println("invalid")
	}else{
		fmt.Println("your age is:",age)
	}
}
