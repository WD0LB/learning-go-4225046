package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")
	states := make(map[string]string, 0)
	states["IT"] = "Information Technology"
	states["HR"] = "Humain Resources"
	states["ADM"] = "Administration"
	fmt.Println(states)
	delete(states, "HR")
	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println("\nsorted")
	sort.Strings(keys)
	for i := range len(keys) {
		fmt.Printf("%v : %v\n", keys[i], states[keys[i]])
	}
}
