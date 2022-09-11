package utils

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter whatever you want")
	input, _ := reader.ReadString('\n')
	fmt.Println("So you've entered ", input)
	fmt.Printf("Type of whatever you've entered is %T", input)
}
