package main

import (
	"fmt"
)

func main() {
	town100 := []string{"Victor", "maa"}

	c := map[string]int{"MMM": 10, "Victor": 123}

	for key, _ := range c {
		if !contains(town100, key) {
			delete(c, key)
		}
	}
	fmt.Println(c)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
