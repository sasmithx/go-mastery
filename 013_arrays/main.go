package main

import (
	"fmt"
)

func main() {
	// fixed and can not grow
	var marks [3]int 

	marks[0] = 85
	marks[1] = 90
	marks[2] = 95

	fmt.Println("Marks:", marks)

	//array literal
	res := [5]int{1,2,3,4,5}
	fmt.Println(len(res))
}