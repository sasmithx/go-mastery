package main

import (
	"fmt"
)

func main() {
	isLoggedIn := true // inferred as boolean type
	isAdmin := false
	hasSubscription := true

	//AND &&
	canOpenDashboard := isLoggedIn && isAdmin

	canDeletePost := isAdmin || (isLoggedIn && hasSubscription)

	fmt.Println(canOpenDashboard, canDeletePost)

	age := 25
	isAdult := age >= 18
	fmt.Println(isAdult)
}