package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его

	"golang.org/x/text/cases"
	// пакет используется для проверки ответа, не удаляйте его
)

func main() {
	var value1, value2, operation interface{} = 12 // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	switch v := value1.(type) {
	case int64:
		fmt.Println("111")
	default:
		fmt.Println("value=полученное_значение: тип_значения")
	}
}
