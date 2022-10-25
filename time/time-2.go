package main

import (
	"fmt"
	"time"
)

func main() {
	first := time.Now()
	second := time.Now().Add(time.Hour * 12)

	fmt.Println(first.Before(second))
	fmt.Println(second.After(first))
	fmt.Println(first.Equal(second))

	fmt.Println(second.Sub(first))
}
