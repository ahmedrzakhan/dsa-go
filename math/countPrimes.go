/**
204. Count Primes

Given an integer n, return the number of prime numbers that are strictly less than n.



Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
Example 2:

Input: n = 0
Output: 0
Example 3:

Input: n = 1
Output: 0


Constraints:

0 <= n <= 5 * 106
*/

package main

import "fmt"

// O(sqrt(n))
func isPrimeCP(N int) bool {
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

func countPrimes(maxNum int) int {
	count := 0
	for i := 1; i < maxNum; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func mainCP() {
	maxNum := 100
	fmt.Println(countPrimes(maxNum))
}
