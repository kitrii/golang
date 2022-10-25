package main

import "fmt"

func main() {
	var num int
	count := 0
	max := 0

	for fmt.Scan(&num); num != 0; fmt.Scan(&num) {
		if num > max {
			max = num
			count = 1
			fmt.Println("num>max", max, count)
		} else if num == max {
			count++
			fmt.Println("num==max", count)
		}
	}
	fmt.Println(max)
	fmt.Println(count)
}
