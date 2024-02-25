package main

import (
	"fmt"
)

/**
A number is considered perfect if its digits sum up to exactly 10.
Given a positive integer n, return the n-th perfect number.
For example, given 1, you should return 19. Given 2, you should return 28.
*/

func digitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10 // Extract Last Digit
		num /= 10       // Remove Last Digit
	}
	return sum
}

func findNthPerfectNumber(n int) int {
	count := 1
	current := 19 // Starting from 19 as the first perfect number

	for count < n {
		current++
		if digitSum(current) == 10 {
			count++
		}
	}
	return current
}

func mainPN() {
	n := 2
	perfectNumber := findNthPerfectNumber(n)
	fmt.Printf("The %d-th perfect number is: %d\n", n, perfectNumber)
}
