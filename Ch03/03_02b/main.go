package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")
	anInt := 10
	var p *int = &anInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value", *p)
	}

	value1 := 12.34
	pointer1 := &value1
	*pointer1 = *pointer1 * 10
	fmt.Println("Pointer", *pointer1)
	fmt.Println("Value", value1)

}
