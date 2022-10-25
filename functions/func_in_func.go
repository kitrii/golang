package main

import "fmt"

func main() {
	type strTypeFunc func(string)

	writer := func(prefix string) strTypeFunc {
		return func(in string) {
			fmt.Println(prefix, in)
		}
	}
	writer("lalala")
}
