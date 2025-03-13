package main

import (
	"fmt"
	"mylearning/function"
	// "mylearning/myinput"
)

func main() {
	fmt.Println("Hello")
	// myinput.CustomInput()

	var version = "1.0"
	fmt.Println(version)

	addition := function.Add(10, 20)
	fmt.Println("Addition of two numbers :", addition)

}
