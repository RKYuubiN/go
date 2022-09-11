package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Conditionals() {
	fmt.Println("All about conditionals")

	loginCount := 12

	var result string

	if loginCount < 10 && loginCount > 1 {
		result = "Regular User"
	} else if loginCount <= 1 {
		result = "New User"
	} else {
		result = "Are you a bot?"
	}

	fmt.Println(result)

	if 6%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// `if err != nil {

	// }`

	fmt.Println("Switch Case")

	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	// Break statement is automatically provided by go
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot")
	default:
		fmt.Println("What???")
	}
}
