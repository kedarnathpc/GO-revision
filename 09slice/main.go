package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices !!!")

	var animals = []string{"cat", "dog", "cow", "sheep"}
	animals = append(animals, "tiger", "lion")
	fmt.Println("The slice is ", animals, "\nand the length of slice is ", len(animals))
	fmt.Printf("The type is %T\n\n", animals)

	animals = append(animals[1:3])
	fmt.Println(animals)

	//memory allocation

	scores := make([]int, 3)
	scores[0] = 8
	scores[1] = 6
	scores[2] = 4
	scores = append(scores, 4, 5, 6)
	fmt.Println(scores)

	sort.Ints(scores)
	fmt.Println(scores)

	//remove values from slice based on index

	var courses = []string{"react", "JS", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
