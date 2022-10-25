package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, ok := ioutil.ReadFile("text.txt")

	if ok != nil {
		panic("Error")
	}
	fmt.Printf("%s", content)

	message := []byte("Message")
	ioutil.WriteFile("write.txt", message, 0600)

}
