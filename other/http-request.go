package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://youtube.com"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Body)
	fmt.Println(response.Status)
	defer response.Body.Close()
}
