package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type FileStructure struct {
	List []struct {
		Global_id int
	}
}

// system_object_id string
// Kod string
// is_deleted string
// signature_date string
// Nomdescr string
// Idx string
// Razdel string
// Name string

func main() {
	text, _ := ioutil.ReadFile("data-20190514T0100.json")
	var t FileStructure

	json.Unmarshal(text, &t)
	sum := 0
	for _, id := range t.List {
		sum = sum + id.Global_id
	}
	fmt.Println(sum)
}
