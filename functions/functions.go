//functions

/*func имя_функции (список_параметров) (типы_возвращаемых_значений) {
    выполняемые_операторы
}
*/

package main

import "fmt"

func main() {
	add(10, 15)
	add(20, 25)
	n, f := client(10, 20, "Ser", "Mir")
	fmt.Println(n)
	fmt.Println(f)
	raspakovka()
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func client(x, y int, name, surname string) (int, string) {
	age := x + y
	fio := fmt.Sprintf("%q --- %q", name, surname)
	return age, fio
}

func sum(args ...int) (summa int) {
	for _, elem := range args {
		summa += elem
	}
	return summa
}

func raspakovka() {
	num := []int{1, 2, 3, 490}
	fmt.Println(sum(num...))

}
