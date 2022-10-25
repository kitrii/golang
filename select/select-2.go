package main

import "fmt"

func main() {
	channel1 := make(chan int, 2)
	channel1 <- 10
	channel1 <- 20

	channel2 := make(chan int, 1)
	channel2 <- 999
LOOP:
	for {
		select {
		case cha := <-channel1:
			fmt.Println("channel1 input:", cha)
		case val := <-channel2:
			fmt.Println("channel2 input", val)
		default:
			break LOOP
		}
	}

}
