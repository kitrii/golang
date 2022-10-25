package main

import "fmt"

func main() {
	defer fmt.Println("the first defer")
	defer fmt.Println("the second defer")
	defer fmt.Println("Само значение инкремента", increment(10))
	fmt.Println("the first string")
	fmt.Println("the second string")

}

func increment(x int) (rez int) {
	fmt.Println("Выполнение функции инкремента")
	rez = x + 1
	return
}
