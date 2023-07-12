package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio, OS

func main() {
	fmt.Println("User Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something : ")

	//comma ok || comma error
	input, _ := reader.ReadString('\n')
	fmt.Printf("Your input is %vand the type is %T\n", input, input)
}
