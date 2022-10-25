package main

import "fmt"

type User struct {
	ID       int
	RealName string `unpack"-"`
	Login    string
	Flags    int
}

func PrintReflect(i interface{}) error {

}

func main() {
	user := &User{
		ID:       18,
		RealName: "Ekaterina",
		Flags:    07122003,
	}

	err := PrintReflect(user)

	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}
