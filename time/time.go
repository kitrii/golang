package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("02-01-2006 15:04:05"))
	future := now.Add(time.Hour * 12)
	fmt.Println(future)
	past := now.AddDate(-10, -10, -7)
	fmt.Println(past)

	current := time.Date(
		2022,
		time.October,
		8,
		12,
		47,
		50,
		54,
		time.UTC,
	)
	fmt.Println(current)

	unixTime := time.Unix(
		2000000000,
		1,
	)
	fmt.Println(unixTime)
	fmt.Println("-----")
	fmt.Println(current.Year())
	fmt.Println(current.Day())
	fmt.Println(current.Hour())
	fmt.Println(current.Second())
	fmt.Println("-----")
	fmt.Println()
}
