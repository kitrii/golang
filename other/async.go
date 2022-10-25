package main

import (
	"fmt"
	"time"
)

func GetComments() chan string {

	result := make(chan string, 1)

	go func(receive chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("func get comments")
		receive <- "32 комментария"

	}(result)
	return result

}

func GetPages() {
	resultCh := GetComments()

	time.Sleep(2 * time.Second)
	fmt.Println("func get pages")

	commentsData := <-resultCh
	fmt.Println("Считали из канала")
	fmt.Println(commentsData)

}

func main() {
	GetPages()
}
