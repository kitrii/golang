package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Adress string
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int {
	return p.Age
}

func (p Person) To80() (rez int) {
	rez = 80 - p.Age
	return
}

func (p *Person) ChangeAge(age int) {
	p.Age = age
}

func main() {
	Me := Person{
		"Ekaterina",
		18,
		"Moscow",
	}
	fmt.Println(Me.Name)
	fmt.Println(Me.Age)
	fmt.Println(Me.Adress)

	Me.ChangeAge(19)
	Me.GetAge()

	fmt.Println(Me.Name)
	fmt.Println(Me.Age)
	fmt.Println(Me.Adress)

}
