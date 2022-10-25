package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name   string
	Age    int
	Weight int
	High   int
}

func main() {

	Ivan := Man{Name: "Ivan", Age: 19, Weight: 78, High: 192}
	marshalling_data, _ := json.Marshal(Ivan)
	fmt.Println(marshalling_data)
	fmt.Printf("%s", marshalling_data)
	fmt.Println()
	fmt.Println(json.Valid(marshalling_data))

	info := `{"Name": "Ivan", "Age": 19, "Weight": 78, "High": 192}`

	var unmarshaling_data Man
	json.Unmarshal([]byte(info), &unmarshaling_data)

	fmt.Println(unmarshaling_data)

}
