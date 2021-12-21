package main

import (
	"fmt"
	"learngo/mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	mypackage.PrintMessage("Hola Por fin lo solucione")
}
