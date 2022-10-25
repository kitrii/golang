package main

import (
	"fmt"
)

func main() {
	a := 5
	defer fmt.Println(add(a))
	a = 6
	fmt.Println(a)

}

func add(num int) int {
	fmt.Println(num)
	num++
	return num
}
