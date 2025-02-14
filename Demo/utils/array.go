package utils

import "fmt"

func ArrayOperation(){

	var arr [5]int // Declares an integer array of size 5

	var arr1 = [3]int{10, 20, 30} //Declare Array of size 3 and initialize values
	arr2 := [3]string{"Go", "Lang", "Array"} // Shorthand notation

	//[...] for auto sizing
	arr3 := [...]int{1, 2, 3, 4, 5} // Compiler determines the size (5)

	fmt.Println(arr)
	fmt.Println(arr1, arr2, arr3)

	fmt.Println("Length of an array:",len(arr3))

	//slice
	slice := []string{"Go", "Lang", "Slice"}

	fmt.Println(slice)

	//meke() 

	slice1 := make([]int,5,10)
	fmt.Println(slice1, len(slice1), cap(slice1))

	var scores =[]int{100, 50, 40}
	scores[2] = 25 // add new value using position

	scores = append(scores, 55) //add new value to the slice

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := arr2[1:3] 

	rangeTwo := arr2[:2]

	rangeOne = append(rangeOne, "operations")

	fmt.Println(rangeOne, rangeTwo)
	

}