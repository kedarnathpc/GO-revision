package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	age := 23
	var vote string

	if age > 18 {
		vote = "You can vote"
	} else if age < 18 {
		vote = "You cannot vote"
	} else {
		vote = "Come next year"
	}

	fmt.Println(vote)

	if num := 45; num > 50 {
		fmt.Println("Number is greater than 50")
	} else {
		fmt.Println("Number is not greater than 50")
	}
}
