package main

import "fmt"

func main() {
	sum, multiplication := calculator(2, 3, 4)

	fmt.Println("Sum is", sum)                       // 9 (2 + 3 + 4)
	fmt.Println("Multiplication is", multiplication) // 24 (2 * 3 * 4)
}

func calculator(a, b, c int) (int, int) {
	s := a + b + c
	m := a * b * c

	return s, m
}
