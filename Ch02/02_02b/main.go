package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	fmt.Println(str1, str2, str3)

	aNumb := 123
	length, err := fmt.Printf("The number is: %v\n", aNumb)
	if err == nil {
		fmt.Println("The length:", length)
	}
	fmt.Printf("The type of number %v is: %T\n", aNumb, aNumb)

}
