package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang !!!")
	greeter()
	sum := adder(2, 5)
	fmt.Println(sum)
	total, message := multiadder(3, 4, 56, 7, 3, 3, 4, 5, 67, 2, 21)
	fmt.Println(total, " ", message)
}

func greeter() {
	fmt.Println("The greeter function is called and executed !!!")
}

func adder(a int, b int) int {
	return a + b
}

func multiadder(val ...int) (int, string) {
	total := 0

	for _, val := range val {
		total = total + val
	}
	return total, "This is done by multiadder"
}
