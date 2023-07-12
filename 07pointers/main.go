package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var ptr *int //create a pointer
	fmt.Println("the value of the pointer is ", ptr)

	num := 33
	var ptr1 = &num
	fmt.Printf("The value of the pointer is %v\nand the value at the pointer is %v\n", ptr1, *ptr1)

	*ptr1 = *ptr1 + 2
	fmt.Println("the new value at pointer is ", *ptr1)
}
