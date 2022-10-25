package main

import "fmt"

func main() {
	defer saveInfo()
	defer fmt.Println("LOLOLOLO")
	fmt.Println("I am just printing something that haven't meaning")
	panicOn(0)
}

func saveInfo() {
	fmt.Println("Save information")
}

func panicOn(num int) {
	if num == 0 {
		panic("O MY Godness, you wrote zero")
	}
}
