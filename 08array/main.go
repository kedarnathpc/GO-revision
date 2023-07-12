package main

import "fmt"

func main() {
	fmt.Println("Welcome to array !!!")

	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "mango"
	fruits[3] = "banana"
	fmt.Println("The fruit array is ", fruits, "\nand the length of array is ", len(fruits))
}
