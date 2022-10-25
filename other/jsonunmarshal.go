package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Markets []struct {
		Name  string
		Price int
	}
}

type status struct {
	Name string
}

func main() {
	text := `{"markets":[{"name":"rainbow", "price":200},{"name":"magnit", "price":10}]}`

	var information info

	json.Unmarshal([]byte(text), &information)
	fmt.Println(information)
	fmt.Printf("%T", information.Markets[0].Price)
	fmt.Printf("%T", information.Markets[1].Price)

}
