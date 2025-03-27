package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	rh := Emp{"RH", 1000}
	fmt.Println(rh)
	rh.Salary = 1200
	fmt.Println(rh)
	fmt.Printf("%+v\n", rh)
}
type Emp struct {
	Post string
	Salary int
}