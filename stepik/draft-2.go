package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scan(&word)
	reverse(&word)
	word2 := reverse(&word)
	fmt.Println(word2)

}

func reverse(word *string) []rune {
	wordRune := []rune(*word)
	lenword := len(wordRune)
	neword := []int32{}

	for i := lenword - 1; i != -1; i-- {
		neword = append(neword, string(wordRune[i]))
	}

	strings.Replace
	return neword
}
