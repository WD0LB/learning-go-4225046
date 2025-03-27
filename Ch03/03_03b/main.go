package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")
	var numb = [3]int{10, 20, 30}
	fmt.Println(numb)

	var colors [2]string
	colors[0] = "White"
	colors[1] = "Black"
	fmt.Println(colors)
	fmt.Println(colors[1])
}
