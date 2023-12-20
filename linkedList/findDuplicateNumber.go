package main

import "fmt"

/**
287. Find the Duplicate Number

Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.

Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
Example 2:

Input: nums = [3,1,3,4,2]
Output: 3

Constraints:

1 <= n <= 105
nums.length == n + 1
1 <= nums[i] <= n
All the integers in nums appear only once except for precisely one integer which appears two or more times.

Follow up:

How can we prove that at least one duplicate number must exist in nums?
Can you solve the problem in linear runtime complexity?
*/

/*
The findDuplicate function works by treating the numbers in the array as pointers in a linked list, then
applying Floyd's Tortoise and Hare cycle detection algorithm. Here's the breakdown:
First Phase (Finding the Intersection): Treat the array as a linked list where each element points to
the index of the next element. slow and fast pointers traverse this "list". Since there's a duplicate,
a cycle is guaranteed. fast moves twice as fast as slow, and they eventually meet inside the cycle.
Second Phase (Finding the Cycle Start): Reset slow to the beginning (index 0). Move both slow and fast
at the same pace. The point where they meet again is the start of the cycle, which corresponds to the
duplicate number.
This works because the duplicate number forms a cycle by pointing to an already visited index, and Floyd's
algorithm efficiently finds the start of this cycle.
*/
// TC - O(N), SC - O(1)
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	// Phase 1: Finding the intersection point of the two runners.
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// Phase 2: Finding the entrance to the cycle.
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func maindup() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println("The duplicate number is:", findDuplicate(nums))
}
