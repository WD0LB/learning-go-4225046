package main

import (
	"fmt"
	"sort"
)
func main() {
	// This is an slice
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	fmt.Printf("type: %T\n", colors)
	var countries = make([]string, 0, 3)
	countries = append(countries, "Morocco", "Egypt", "India")
	fmt.Println(countries)
	countries = remove(countries, 2)
	fmt.Println(countries)
	sort.Strings(countries)
	fmt.Println(countries)
}
func remove(slice []string, index int) []string {
	return append(slice[:index],slice[index+1:]...)
}
