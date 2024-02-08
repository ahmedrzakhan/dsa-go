package main

import "fmt"

// O(sqrt(n))
func isPrime(N int) bool {
	if N <= 1 {
		return false // Numbers less than or equal to 1 are not prime
	}
	if N <= 3 {
		return true // 2 and 3 are prime
	}

	// Check if n is divisible by 2 or 3
	if N%2 == 0 || N%3 == 0 {
		return false
	}

	for i := 5; i*i <= N; i += 6 {
		if N%i == 0 || N%(i+2) == 0 {
			return false
		}
	}

	return true
}

func sieveOfEratosthenes(maxNum int) []bool {
	// Create a boolean slice, where prime[i] will be true if i is prime
	primes := make([]bool, maxNum+1)

	for i := 2; i <= maxNum; i++ {
		if isPrime(i) {
			primes[i] = true
		} else {
			primes[i] = false
		}
	}

	return primes
}

func mainSE() {
	maxNum := 100
	primeMarks := sieveOfEratosthenes(maxNum)

	// Printing prime numbers and their truth value
	for i := 1; i <= maxNum; i++ {
		fmt.Printf("Number %d is prime: %v\n", i, primeMarks[i])
	}
}
