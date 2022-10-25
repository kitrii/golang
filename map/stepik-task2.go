package main

import "fmt"

func main() {
	var friendsOfDima []string
	friendsOfSymon := make([]string, 3)
	fmt.Printf("%T", friendsOfDima)
	fmt.Printf("%T", friendsOfSymon)
	friendsOfDima = append(friendsOfDima, "Костя", "Semen", "Tanya")
	friendsOfSymon = append(friendsOfSymon, "Valera", "Tanya", "Dima")
	fmt.Println()
	fmt.Println(friendsOfDima)
	fmt.Println(friendsOfSymon)
	friends := map[string][]string{
		"Dima":   friendsOfDima,
		"Symon":  friendsOfSymon,
		"Kostya": nil,
	}
	fmt.Println(friends)
	key, value := friends["Kostya"]
	fmt.Println(key, value)
	fmt.Println(friendsOfSymon[0])

	fmt.Println(friendsOfSymon[1])

	fmt.Println(friendsOfSymon[3])

}
