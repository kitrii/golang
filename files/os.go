package main

import (
	"fmt"
	"os"
)

func main() {
	// Функции пакета os
	// func Create(fileName string) (*Func, err) создание файла с именем name
	// func Open(fileName string) (*File, error) // открытие файла с именем name
	file, _ := os.Create("numbers.txt")
	fmt.Println(file.Name())
	os.Rename("numbers.txt", "new_numbers.txt")
	info, _ := os.Stat("numbers.txt")
	fmt.Println(info)
	file.WriteString("Kadabra")
	file.WriteString("Something")

	fmt.Println(file.Name())
	os.Open(file.Name())

}
