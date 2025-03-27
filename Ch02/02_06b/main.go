package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")
	t := time.Date(2025, time.March, 16, 48, 0, 0, 0, time.UTC)
	fmt.Println("The time %s", t)

}
