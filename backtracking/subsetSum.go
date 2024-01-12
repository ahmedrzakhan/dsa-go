package main

import "fmt"

/**
Subset Sum Problem

https://www.geeksforgeeks.org/problems/subset-sum-problem-1611555638/1

Given an array of non-negative integers, and a value sum, determine if there is a subset
of the given set with sum equal to given sum.

Example 1:

Input:
N = 6
arr[] = {3, 34, 4, 12, 5, 2}
sum = 9
Output: 1
Explanation: Here there exists a subset with
sum = 9, 4+3+2 = 9.
Example 2:

Input:
N = 6
arr[] = {3, 34, 4, 12, 5, 2}
sum = 30
Output: 0
Explanation: There is no subset with sum 30.

Your Task:
You don't need to read input or print anything. Your task is to complete the function isSubsetSum() which takes the array arr[], its size N and an integer sum as input parameters and returns boolean value true if there exists a subset with given sum and false otherwise.
The driver code itself prints 1, if returned value is true and prints 0 if returned value is false.

Expected Time Complexity: O(sum*N)
Expected Auxiliary Space: O(sum*N)

Constraints:
1 <= N <= 100
1<= arr[i] <= 100
1<= sum <= 104
*/

func isSubsetSum(arr []int, target int) bool {
	return helperSubsetSum(0, arr, target, 0)
}

func helperSubsetSum(i int, arr []int, target int, curSum int) bool {
	if curSum == target {
		return true
	}

	if curSum > target || i >= len(arr) {
		return false
	}

	// Recursive call including the current element
	if helperSubsetSum(i+1, arr, target, curSum+arr[i]) {
		return true
	}

	// Recursive call excluding the current element
	return helperSubsetSum(i+1, arr, target, curSum)
}

func mainSubsetSum() {
	arr := []int{3, 34, 4, 12, 5, 2}
	target := 9

	if isSubsetSum(arr, target) {
		fmt.Println("Found a subset with given sum")
	} else {
		fmt.Println("No subset with given sum")
	}
}
