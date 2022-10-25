package main

import "fmt"

func greet(c chan string) {
	fmt.Println("func greet started")
	fmt.Println("c equal", c)
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	fmt.Println(c, "new channel")

	go greet(c)
	c <- "John"

	fmt.Println("main() stopped")
}
