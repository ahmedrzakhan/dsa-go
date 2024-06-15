package main

import (
	"fmt"
	"sort"
)

/**
find next smallest combination
2187615 -> 251678
*/

// Check Descending Order: If the digits are in descending order, return "Not Possible".
// Find Swap Point: Iterate from right to left until you find a digit digits[i] smaller than digits[i + 1].
// Find Replacement: Find the smallest digit digits[j] on the right of digits[i] where digits[j] > digits[i].
// Swap: Swap digits[i] and digits[j].
// Sort: Sort the subarray from digits[i + 1] to the end.

func findNextSmallestCombination(num string) string {
	digits := []byte(num) // Convert string to byte slice for easier manipulation
	N := len(digits)

	// Find the swap point (first digit from right that is smaller than its next digit)
	swapPoint := -1
	for i := N - 2; i >= 0; i-- {
		if digits[i] < digits[i+1] {
			swapPoint = i
			break
		}
	}

	// If no swap point is found, digits are in descending order (no next combination exists)
	if swapPoint == -1 {
		return "Not Possible"
	}

	// Find the next smallest digit on the right of swapPoint (just larger than digits[swapPoint])
	nextSmallestIndex := -1
	for i := N - 1; i > swapPoint; i-- {
		if digits[i] > digits[swapPoint] {
			nextSmallestIndex = i
			break
		}
	}

	// Swap the swapPoint digit with the nextSmallest digit
	digits[swapPoint], digits[nextSmallestIndex] = digits[nextSmallestIndex], digits[swapPoint]

	// Sort the digits after swapPoint in ascending order
	sort.Slice(digits[swapPoint+1:], func(i, j int) bool {
		return digits[swapPoint+1+i] < digits[swapPoint+1+j]
	})

	// Convert back to string
	return string(digits)
}

func mainN() {
	num := "258761" // 261788
	// num := "258761"
	nextSmallest := findNextSmallestCombination(num)
	fmt.Println("Next smallest combination:", nextSmallest)
}

// 218769 298761
