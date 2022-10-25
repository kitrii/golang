package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	go say("Hey")
	say("There")
}
