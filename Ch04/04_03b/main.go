package main

import "fmt"

func main() {

	letters := []string{"A", "B", "C"}
	for i := 0; i < len(letters); i++ {
		println(letters[i])
	}
	for i := range letters {
		println(letters[i])
	}
	for _, letter := range letters {
		println(letter)
	}

	departements := make(map[string]string)
	departements["IT"] = "Inf Tech"
	departements["HR"] = "Humain Resource"
	departements["ADM"] = "Administration"
	for departement, _ := range departements {
		println(departements[departement])
	}

	numb := 4
	prod := 1
	for prod < 400 {
		prod *= numb
		if prod > 80 {
			goto message
		}
	}
message:
	fmt.Printf("The final prod %v\n", prod)
}
