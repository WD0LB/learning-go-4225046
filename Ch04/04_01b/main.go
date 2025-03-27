package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")
	theAnswer := -1
	var result string
	if theAnswer > 0 {
		result = "Greater than zero"
	} else if theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equals zero"
	}

	fmt.Printf("The results: %s\n", result)
}
