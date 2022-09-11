package main

import "fmt"

func printHello(ch chan int) {
	fmt.Println("hello fromt printHello function")
	ch <- 3
}

func main() {
	ch := make(chan int)
	fmt.Println("using context pacakge in go")
	go func() {
		fmt.Println("Helloo")
		ch <- 2
	}()
	fmt.Println("Calling between direct go and printHello function")
	go printHello(ch)
	go func() {
		fmt.Println("In")
		ch <- 4
	}()
	i := <-ch
	i = <-ch
	fmt.Println("Value of i is ", i)
	i = <-ch
}
