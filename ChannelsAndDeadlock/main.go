package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	// var wg *sync.WaitGroup = &sync.WaitGroup{}
	fmt.Println("Channels in golang")

	myChannel := make(chan int, 1)

	// myChannel <- 5

	// fmt.Println(<-myChannel)
	wg.Add(2)
	// send only goroutine
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(val, isChannelOpen)
		fmt.Println(<-ch)
		// fmt.Println(<-ch)
		wg.Done()
	}(myChannel, wg)

	// receive only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		// ch <- 2
		// ch <- 6
		close(ch)
		wg.Done()
	}(myChannel, wg)
	wg.Wait()
}
