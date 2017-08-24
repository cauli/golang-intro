package main

import "fmt"

func main() {
	sum, multiplication := calculator(2, 3)

	fmt.Println("Sum is", sum)                       // 5
	fmt.Println("Multiplication is", multiplication) // 6
}

func calculator(a, b int) (int, int) {
	s := a + b
	m := a * b

	return s, m
}
