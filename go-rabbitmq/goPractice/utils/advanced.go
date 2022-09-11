package utils

import "fmt"

func Defer() {
	defer fmt.Println("hello")
	defer fmt.Println("hey")
	defer fmt.Println("hi")
	fmt.Println("This is defer")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
