package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 1) //буферизированный канал, заблочится и передаст управление другой горутине только когда будет 2 элемента в ней

	go func(in chan int) {
		fmt.Println("GO: before perceive info")
		val := <-channel
		val2 := <-channel
		fmt.Println("GO: get some info from channel", val, val2)
		fmt.Println("GO: after read from channel")
	}(channel)

	fmt.Println("MAIN: Before to give the info")
	channel <- 134559
	channel <- 134
	fmt.Println("MAIN: the information is in the channel yet")
	fmt.Scanln()
}
