package main

import "fmt"

func main() {
	var slice1 []int64 //slice initializaition
	var slice2 []int64 = []int64{1, 2, 3, 4, 5}

	slice3 := []int64{3, 6, 9, 12, 15}
	slice4 := []int64{10: 100, 1: 1}

	/*make([]T, length, capacity)
		длина (length) — это количество элементов среза;
	    емкость (capacity) - количество элементов между началом среза и концом базового массива.*/
	slice5 := make([]int64, 10, 20)

	fmt.Println(slice1)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(len(slice2)) //slice length
	fmt.Println(cap(slice2)) //slice capacity

	fmt.Println("Оператор среза") //Оператор среза
	fmt.Println(slice3[0:2])

	//copy (idk)
	a := make([]int64, 5, 5)
	b := copy(a, slice2)
	fmt.Println(slice2)
	fmt.Println(b)

	//append
	slice1 = append(slice1, 1, 2, 3, 4)
	fmt.Println("New slice1", slice1)

	slice6 := append(slice1[0:1], slice1[3])
	fmt.Println(slice6)

}
