package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(1 + time.Second)

	select {
	case <-timer.C:
		fmt.Println("Произошло чтение из канала Timer")
	}
}
