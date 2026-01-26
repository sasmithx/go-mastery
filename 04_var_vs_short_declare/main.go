package main

import (
	"fmt"
)

func main() {
	var city string
	city = "New York"

	var name = "Sasmith" // inferred to string

	// :=

	subscribers := 5000 // inferred to int

	subscribers = subscribers + 1000
	
	likes, comments := 300, "25" // inferred to int

	fmt.Println(city, name, subscribers, likes, comments)	
}