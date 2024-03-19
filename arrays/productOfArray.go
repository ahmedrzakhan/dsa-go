package main

import "fmt"

/**
238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal
to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/

// TC - O(N), SC - O(N)
func productExceptSelf(A []int) []int {
	N := len(A)
	result := make([]int, N)

	// Calculate prefix products (products to the left)
	prefix := 1
	for i := 0; i < N; i++ {
		result[i] = prefix
		prefix *= A[i]
	}

	// Calculate suffix products (products to the right)
	postfix := 1
	for i := N - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= A[i]
	}

	return result
}

func mainPA() {
	nums := []int{1, 2, 3, 4}
	productArray := productExceptSelf(nums)
	fmt.Println("Product of Array Except Self:", productArray)
}
