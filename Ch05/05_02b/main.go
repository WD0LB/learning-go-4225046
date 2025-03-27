package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	fmt.Printf("The %v says %v!\n", dog.Breed, dog.Sound)
	fmt.Println(dog.Speak())
}

type Dog struct {
   Breed string
   Sound string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("dog says %v", d.Sound)
}