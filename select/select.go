package main

import "fmt"

func main() {
	channel1 := make(chan int, 1)
	channel2 := make(chan string, 1)

	channel1 <- 123437878
	channel2 <- "12343"
	select {
	case cha := <-channel1:
		fmt.Println(cha)
	case val := <-channel2:
		fmt.Println(val)
	default:
		fmt.Println("Something wrong with channels")
	}
}
