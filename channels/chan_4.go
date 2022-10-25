package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("before add", i)
			out <- i
			fmt.Println("after add", i)

		}
		close(out)
		fmt.Println("OUT is finished")

	}(channel)
	fmt.Println("MAIN")
	for j := range channel {
		fmt.Println("\tget", j)
	}
	//fmt.Scan()
}
