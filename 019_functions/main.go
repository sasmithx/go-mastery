package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func SumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product
}

func main() {
	sum1 := add(3, 5)
	s, p := SumAndProduct(4, 6)
	fmt.Println("Sum1:", sum1, s, p)

	// _

	onlySum, _ := SumAndProduct(10, 2)


	fmt.Println("Sum1:", sum1, s, p)
	fmt.Println("Only Sum:", onlySum)
}