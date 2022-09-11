package utils

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func Conversion() {
	// declare a variable and assign its value to standard input
	randomNumber := bufio.NewReader(os.Stdin)

	// Read the input buffer till it encounters a new line (\n) in this case and store it in input variable
	input, _ := randomNumber.ReadString('\n')
	fmt.Println("Your number is", input)
	increasedNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your increased number is ", increasedNum+1)
	}
	fmt.Printf("Type of random number is %T", randomNumber)

}

func HandlingTime() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 16, 04, 20, 12, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
}
