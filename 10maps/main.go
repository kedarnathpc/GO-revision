package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[int]string)
	languages[0] = "javascript" // adding values to map
	languages[1] = "golang"
	languages[2] = "python"

	fmt.Println("The map is : ", languages)
	fmt.Println("The language at 0 is ", languages[0])

	// deleting values from the map
	delete(languages, 2)
	fmt.Println(languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("The key and their values are %v : %v\n", key, value)
	}
}
