package main

import "fmt"

// struct groups related fields into one type

type User struct{
	ID int
	Name string
	Email string
	Age int
}

func main(){
	u1 := User{
		ID: 1,
		Name: "Sasmith",
		Email: "sasmith@gmail.com",
		Age: 100,
	}

	fmt.Println(u1, u1.ID, u1.Email)

	// mutable by default 
	u1.Age = 400

	fmt.Println(u1)

	u2 := User{
		Name: "jhon",
		Email: "jhon@gmail.com",
	}

	fmt.Println("Partial User", u2)

}