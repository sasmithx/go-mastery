package main

import (
	"fmt"
)

func main() {
	views1 := 1000
	views2 := 2000
	totalViews := views1 + views2

	likes := 500
	likes++
	likes++

	avgViews := totalViews / 2

	fmt.Println(totalViews, likes, avgViews)
}