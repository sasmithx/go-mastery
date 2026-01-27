package main

import (
	"fmt"
)

func main() {
	// common collection type
	//dynamics and can grow
	// []type{...}

	results := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Results:", results, results[0], len(results), results[len(results)-1])

	results[1] = "Jhon"
	fmt.Println("Results:", results)

	var nums []int
	
	nums = append(nums, 10)
	nums = append(nums, 20, 30)

	fmt.Println("Nums:", nums)
}