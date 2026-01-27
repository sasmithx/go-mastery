package main

import (
	"fmt"
)

func main() {

	// make([], len, cap)
	scores := make([]int,0,5)

	fmt.Println("Scores:", scores, "len:", len(scores), "cap:", cap(scores))

	scores = append(scores, 100)
	fmt.Println("after appending 100", scores)

	scores = append(scores, 200, 300)
	fmt.Println("after appending 200, 300", scores)


	scores = append(scores, 45, 50)
	fmt.Println("after appending 45, 50", scores)
	fmt.Println("len:", len(scores), "cap:", cap(scores))

	// if in case we r exceeding capacity, go grows the backing array (usually doubles)
	scores = append(scores, 60)
	fmt.Println("after appending 600", scores)
	fmt.Println("len:", len(scores), "cap:", cap(scores))

	todos := []string{"task1", "task2", "task3"}

	more := []string{"learn go", "learn docker"}

	todos = append(todos, more...)
	fmt.Println("Todos:", todos)
}