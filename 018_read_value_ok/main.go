package main

import (
	"fmt"
)

func main() {
	points := map[string]int{
		"alice": 10,
		"bob":   0, // valid value
	}

	fmt.Println("Alice's points:", points["alice"])
	fmt.Println("Bob's points:", points["bob"])
	fmt.Println("Charlie's points:", points["charlie"]) // not present, returns zero value

	valB, okB := points["bob"]
	fmt.Println(valB, okB)

	valC, okC := points["charlie"]
	fmt.Println(valC, okC)

	if val, ok := points["charlie"]; ok {
		fmt.Println("Charlie's points:", val)	
	} else {
		fmt.Println("Charlie not found")
	}

	prices := map[string]int{
		"apple":  100,
		"banana": 500,
	}

	total := 0
	for item, price := range prices {
		fmt.Println(item, price)
		total += price
	}

	fmt.Println("Total price:", total)

	for item := range prices {
		fmt.Println("Item:", item)
	}
}