package main

import (
	"fmt"

	"github.com/rkyubin/gopractice/utils"
)

func main() {
	fmt.Println("Working with packages")

	x, y := utils.ReturnsTwoValues(2, 3)

	fmt.Println(x, y)

	utils.UserInput()

}
