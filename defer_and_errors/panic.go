package main

import "fmt"

func main() {
	deferTest()
}

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("The recovering, regaining control, there was", err)
		}
	}()
	fmt.Println("Some useful work")
	panic("something bad happend")
}
