package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello world")
	tripleArgument(1, 2, "Hola mi gente")
	value := returnValue(10)
	fmt.Println("Value:", value)
	value1, _ := doubleReturn(20)
	fmt.Println("Value 1:", value1)
}
