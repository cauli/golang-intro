package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	age := 28 // Type inferred (int)
	fmt.Println(age)

	var tooOld int
	tooOld = age * 2
	fmt.Println(tooOld)
}
