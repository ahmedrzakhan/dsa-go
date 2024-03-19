package main

import (
	"fmt"
)

/**
347. Top K Frequent Elements

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.
*/

/**
using max heap
TC - O(N log K), SC - O(N) + O(K)
*/

// TC - O(N), SC - O(N)
func topKFrequent(A []int, K int) []int {
	freqMap := make(map[int]int) // Map to store element frequencies
	maxFreq := 0                 // Track the maximum frequency

	// Count element frequencies
	for _, num := range A {
		freqMap[num]++
		if freqMap[num] > maxFreq {
			maxFreq = freqMap[num]
		}
	}

	// Create buckets based on frequency
	// NOTE: not creating bucket with size of len(A) + 1, more memory efficient

	buckets := make([][]int, maxFreq+1)
	for num, count := range freqMap {
		buckets[count] = append(buckets[count], num)
	}

	// Extract top k elements from buckets
	result := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
		if len(result) == K {
			break // We have gathered the top k elements
		}
	}

	return result
}

func mainTo() {
	A := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(A, k)
	fmt.Println(result)

	A = []int{1, 2, 3, 4, 5}
	k = 5
	result = topKFrequent(A, k)
	fmt.Println(result)
}
