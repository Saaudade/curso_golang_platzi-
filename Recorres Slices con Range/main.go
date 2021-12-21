package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var t = strings.ToLower(text)
	var textReverse string

	for i := len(t) - 1; i >= 0; i-- {
		textReverse += string(t[i])
	}

	if t == textReverse {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i := range slice {
		fmt.Println(i)
	}

	// ama
	// amor a roma

	isPalindromo("Ama")
}
