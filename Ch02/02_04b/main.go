package main

import "fmt"

func main() {
	i1, i2, i3 := 10, 20, 30
	intSum := i1 + i2 + i3
	fmt.Println("Sum is:", intSum)

	f1, f2, f3 := 10.1, 20.2, 30.3
	floatSum := f1 + f2 + f3
	fmt.Println("Sum is:", floatSum)

	sumSum := float64(intSum) + floatSum
	fmt.Println("Sum of sum:", sumSum)
	
}
