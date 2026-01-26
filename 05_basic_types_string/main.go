package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)

	fmt.Println(strings.ToUpper(fullName))
}