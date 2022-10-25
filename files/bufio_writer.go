package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("File.txt")

	defer file.Close()

	writter := bufio.NewWriter(file)

	n, _ := writter.WriteString("Dear friend I would like to tell you...")
	fmt.Printf("%d", n)
	writter.Flush()

}
