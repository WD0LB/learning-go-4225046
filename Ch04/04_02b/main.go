package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	switch dayNumber {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekdays")
	default:
		fmt.Println("Weekends")
	}

}
