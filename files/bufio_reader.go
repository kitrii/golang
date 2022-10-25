package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("text.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	buffer := make([]byte, 100)

	reader.Read(buffer)

	fmt.Printf("%s", buffer)

}
