package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jóse"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["Jóse"]
	fmt.Println(value, ok)
}