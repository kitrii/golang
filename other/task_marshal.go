package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Rating []int
}
type Group struct {
	Students []Student
}
type Rating struct {
	Average float64
}

func main() {

	//var group Group
	count_students := 0.0
	count_marks := 0.0

	text, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(text)

	var m Group

	json.Unmarshal(text, &m)

	for _, mean := range m.Students {
		count_students += 1
		for _, mean2 := range mean.Rating {
			mean2 = mean2
			count_marks += 1
		}
	}

	i := Rating{count_marks / count_students}
	data, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("%s", data)

}
