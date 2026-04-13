package main

import "fmt"

func divide(a int, b int) (jhon int, sasmith int) {
	jhon = a / b
	sasmith = a + b

	return
}

func main() {

	q, r := divide(10, 10)
	fmt.Println(q, r)
}