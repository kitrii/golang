package main

import (
	"fmt"
	"time"
)

func main() {
	var date string
	fmt.Scan(&date)

	newdate := time.Parse("02-01-2006 15:04:05", date)
	

}
