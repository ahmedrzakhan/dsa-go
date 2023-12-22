package main

import "fmt"

/**
875. Koko Eating Bananas

Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile
has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.

Example 1:

Input: piles = [3,6,7,11], h = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23

Constraints:

1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109
*/

/*
NOTE: The key is in the speed - 1 addition. This increment ensures that unless banana is an
exact multiple of speed, the division rounds up. If banana is an exact multiple of speed,
adding speed- 1 doesn't change the outcome of the division after rounding down.
*/

func canEat(piles []int, hours, speed int) bool {
	timeTaken := 0
	for _, banana := range piles {
		timeTaken += (banana + speed - 1) / speed
		if timeTaken > hours {
			return false
		}
	}

	return true
}

func maxInArray(arr []int) int {
	max := arr[0]
	for _, ele := range arr {
		if ele > max {
			max = ele
		}
	}
	return max
}

// brute : O(max(P(P)))
// TC - O(Log(max(P)P), S - O(1)
func minEatingSpeed(piles []int, h int) int {
	L, H, res := 1, maxInArray(piles), 1
	for L <= H {
		mid := (L + H) / 2
		if canEat(piles, h, mid) {
			res = mid
			H = mid - 1
		} else {
			L = mid + 1
		}
	}

	return res
}

func mainBan() {
	arr := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(minEatingSpeed(arr, h))
}
