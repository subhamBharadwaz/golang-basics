package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 5, 8, 5)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro message is: ", myMessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
