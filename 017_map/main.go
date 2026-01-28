package main

import (
	"fmt"
)

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"Jhon":  25,
		"Maria": 30,
		"Bob":   22,
	}

	fmt.Println("age map:", ages["Maria"], len(ages))

	// make (map[K]V)

	var scores map[string]int // nil map

	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)

	scores["math"] = 95

	fmt.Println("scores:", scores, scores["math"])

	users := map[string]string{
		"u1": "Jhon",
		"u2": "Maria",
		"u3": "Bob",
	}

	fmt.Println(users)

	delete(users, "u2")
	delete(users, "u100") // no error

	fmt.Println("after delete:", users)

}