// massive
package main

import "fmt"

func main() {
	//Объявление массива
	print("Объявление массива\n")

	var m1 [5]int64
	fmt.Println(m1)

	var m2 [10]int64 = [10]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(m2)

	var m3 [5]int64 = [5]int64{1: 1, 4: 4}
	fmt.Println(m3)

	var m4 = [...]int64{0, 3, 6, 9}
	fmt.Println(m4)

	// Обращение ко всем элементам - цикл
	print("Обращение ко всем элементам - цикл\n")
	for i := 0; i < len(m4); i++ {
		fmt.Println(m4[i])
	}

	// Обращение ко всем жлементам через цикл и range
	print("Обращение ко всем элементам - цикл и range\n")
	for index, elem := range m2 {
		fmt.Println(index, elem)
	}

	for ind := range m2 {
		fmt.Println(m2[ind])
	}

}
