package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Job  string
}

type Place struct {
	Id           int
	WorkingYears func(int) int
	Owner        Person
}

func main() {
	var objectPlace = Place{
		Id:    1,
		Owner: Person{"Michael", 25, "IT-spesialist"},
	}
	objectPlace.WorkingYears()
	fmt.Println(objectPlace)
}
