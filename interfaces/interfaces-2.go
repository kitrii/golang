package main

import "fmt"

func main() {
	x := 10
	Add(&x)
	fmt.Println("Переменная в main")
	fmt.Println(x)
	fmt.Println(&x)
}

func Add(x *int) {
	fmt.Println("Переменная в функции")
	*x += 1
	fmt.Println(*x)
	fmt.Println(&x)
}
