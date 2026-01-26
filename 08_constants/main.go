package main

import (
	"fmt"
)

func main() {
	const Pi = 3.14
	const AppName = "GoApp"
	const MaxUsers = 1000

	// typed constants
	
	const Version string = "1.0.0"

	const discountRate float64 = 0.1

	fmt.Println("App Name:", AppName, "Version:", Version, "Pi:", Pi, "Max Users:", MaxUsers, "Discount Rate:", discountRate)
}