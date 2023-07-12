package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Golang")

	//a dice roll game
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open")
	case 2:
		fmt.Println("You can move two spots")
		// fallthrough
	case 3:
		fmt.Println("You can move three spots")
	case 4:
		fmt.Println("You can move four spots")
	case 5:
		fmt.Println("You cam move five spots")
	case 6:
		fmt.Println("You can move six spots and roll the dice again")
	default:
		fmt.Println("What was that !")
	}
}
