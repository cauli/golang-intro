package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Cannot change size of array!
	// Slices to the rescue

	var firstThreePrimes []int
	firstThreePrimes = primes[0:3]

	sum := sumInSlice(firstThreePrimes)
	fmt.Println("Slice is...", firstThreePrimes)
	fmt.Println("Sum is of first three... ", sum)

	// This is how we add to a slice
	firstThreePrimes = append(firstThreePrimes, 15)
	fmt.Println("Sum is of first three + 15... ", sumInSlice(firstThreePrimes))

	// But if we try to add to an array we... just can't
	// primes = append(primes, 55)
}

func sumInSlice(slice []int) int {
	sum := 0

	for _, prime := range slice {
		sum += prime
	}

	return sum
}
