package main

import "fmt"

func main() {
	anonimFunc := func(n1, n2 int) {
		fmt.Println(n1 + n2)
	}

	func(in string) {
		fmt.Println("anonimus string:", in)
	}("lalalalalalala")

	anonimFunc(10, 20)
	return

	type an func(int)
	multiply := func(variable an) {
		variable(120)
		 
	}
	
	return


}
