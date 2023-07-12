package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	fmt.Printf("Please enter rating for service : ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	//conversion

	plusrating, er := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("added one to your rating : ", plusrating+1)
		fmt.Printf("the type is %T: ", plusrating)
	}
}
