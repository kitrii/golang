package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		text := in.Text()
		if prev == text {
			continue
		} else if prev > text {
			return fmt.Errorf("Файл не отсортирован")
		} else {
			prev = text
			fmt.Fprintln(output, text)
		}
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic("Файл не отсортирован")
	}

}
