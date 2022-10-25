package main

import (
	"fmt"
)

const (
	iterationNum = 7
	goroutineNum = 5
)

func doSomeWork(integer int) {
	for j := 0; j < integer; j++ {
		fmt.Println(integer)
	}
}

func main() {
	for i := 0; i < iterationNum; i++ {
		doSomeWork(i)
	}
}
