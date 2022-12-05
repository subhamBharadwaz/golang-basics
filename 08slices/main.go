package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice in golang")

	var fruitList = []string{"Apple", "Banana", "Orange"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 965
	highScores[2] = 465
	highScores[3] = 865
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	// how to remove a value from slice based on index

	var courses = []string{"react", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted((highScores)))
}
