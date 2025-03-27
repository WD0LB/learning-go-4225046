package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	world := "World"
	greeting, count := helloName(world, "Me", "You")
	fmt.Println("Greeting:",greeting)
	fmt.Println("Counting:", count)
}

func helloName(names ...string) (string, int) {
	greeting := "Hello "
	for _, name := range names {
		greeting += name + " "
	}
	count := len(names)
	return greeting, count
}