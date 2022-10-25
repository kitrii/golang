package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var prev string
	for input.Scan() {
		text := input.Text()
		if prev == text {
			continue
		} else if prev > text {
			panic("Файл не отсортирован")
		} else {
			prev = text
			fmt.Println(text)
		}

	}
}
