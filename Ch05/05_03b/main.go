package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	go hi("World")
	fmt.Println("Let's go")
	time.Sleep(2 * time.Second)
	fmt.Println("End")

}

func hi(name string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hi", name)
}
