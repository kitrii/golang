package main

import "fmt"

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	groupCity := map[int][]string{
		10:  []string{"No"},             // города с населением 10-99 тыс. человек
		100: []string{"Moscow", "Novo"}, // города с населением 100-999 тыс. человек
	}

	cityPopulation := map[string]int{
		"Moscow": 200,
		"MMMMM":  12300,
	}

	towns100 := groupCity[100]

	for city, _ := range cityPopulation {
		if !Contains(towns100, city) {
			delete(cityPopulation, city)
		}
	}
	fmt.Println(cityPopulation)
}
