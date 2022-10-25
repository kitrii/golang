package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("%T", a)
	fmt.Println(add(a, b))
}

func add(num1, num2 int) int {
	c := num1*num1 + num2*num2
	return c
}
