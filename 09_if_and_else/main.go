package main

import (
	"fmt"
)

func main() {
	score := 78

	if score >= 75 {
		fmt.Println("A Grade")
	}else if score >= 60 {
		fmt.Println("B Grade")
	}else if score >= 50 {
		fmt.Println("C Grade")
	} else {
		fmt.Println("D Grade")
	}
}