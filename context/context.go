package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	processLongTime()

	ctx := context.Background()
	fmt.Println(ctx)

	_, cancel := context.WithCancel(ctx)

	fmt.Println(cancel)

}
func processLongTime() {
	time.Sleep(0 * time.Second)
	fmt.Println("I wait 2 seconds")
}
