package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance i golang; no super no parent

	subham := User{"Subham", "subham@go.dev", true, 23}
	fmt.Println(subham)
	fmt.Printf("subham details are: %+v\n", subham)
	fmt.Printf("Name is %+v and email is %v.", subham.Name, subham.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
