package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	alreadySeen := make(map[string]bool)
	for input.Scan() {
		text := input.Text()
		// fmt.Println(text)
		// fmt.Println(alreadySeen[text])
		if _, found := alreadySeen[text]; found {
			continue
		}
		alreadySeen[text] = true
		fmt.Println(text)
	}
	fmt.Println(alreadySeen)
}
