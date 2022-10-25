package main

import "fmt"

func main() {
	var num int64 = 12
	num2 := uint16(num)
	fmt.Println(num2)
}

var summ int = 0

func sum(num int) int {

	if num < 10 {
		fmt.Println("if num =", num)
		summ += num
		fmt.Println("if sum", summ)
		return summ

	} else {
		fmt.Println("else num = ", num)
		fmt.Println("num%10 =", num%10)
		summ += num % 10
		fmt.Println("else summ", summ)
		return sum(sum(num / 10))
	}

}
