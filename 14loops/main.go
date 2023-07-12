package main

import "fmt"

func main() {
	days := []string{"sun", "mon", "tues", "wed", "thurs", "fri", "sat"}

	//for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//range based
	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("The index is %v, and day is %v\n", index, day)
	}

	//while loop
	i := 1
	for i < 10 {
		if i == 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}
