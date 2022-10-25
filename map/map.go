package main

import "fmt"

func main() {
	m := make(map[string]int)
	n := map[int]int{1: 2, 3: 4}
	fmt.Println(m)
	fmt.Println(n[1])
	value1, ok2 := n[3]
	value, ok := n[4]
	fmt.Println(value1, ok2)
	fmt.Println(value, ok)

	if value, ok := n[1]; ok == true {
		fmt.Println(value)
	} else {
		fmt.Println("This varuable is not exist")
	}
}
