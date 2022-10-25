package main

import (
	"fmt"
)

func main() {
	channel := make(chan int) //небуферизированный канал

	go func(in chan int) {
		fmt.Println("GO: before perceive info")
		val := <-channel
		fmt.Println("GO: get some info from channel", val)
		fmt.Println("GO: after read from channel")
	}(channel)

	fmt.Println("MAIN: Before to give the info")
	channel <- 134559
	fmt.Println("MAIN: the information is in the channel yet")
}
