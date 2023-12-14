package main

import (
	"fmt"
	"strconv"
)

/**
You are keeping the scores for a baseball game with strange rules. At the beginning of the game, you start with an empty record.

You are given a list of strings operations, where operations[i] is the ith operation you must apply to the record and is one of the following:

An integer x.
Record a new score of x.
'+'.
Record a new score that is the sum of the previous two scores.
'D'.
Record a new score that is the double of the previous score.
'C'.
Invalidate the previous score, removing it from the record.
Return the sum of all the scores on the record after applying all the operations.

The test cases are generated such that the answer and all intermediate calculations fit in a 32-bit integer and that all operations are valid.



Example 1:

Input: ops = ["5","2","C","D","+"]
Output: 30
Explanation:
"5" - Add 5 to the record, record is now [5].
"2" - Add 2 to the record, record is now [5, 2].
"C" - Invalidate and remove the previous score, record is now [5].
"D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
"+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
The total sum is 5 + 10 + 15 = 30.
Example 2:

Input: ops = ["5","-2","4","C","D","9","+","+"]
Output: 27
Explanation:
"5" - Add 5 to the record, record is now [5].
"-2" - Add -2 to the record, record is now [5, -2].
"4" - Add 4 to the record, record is now [5, -2, 4].
"C" - Invalidate and remove the previous score, record is now [5, -2].
"D" - Add 2 * -2 = -4 to the record, record is now [5, -2, -4].
"9" - Add 9 to the record, record is now [5, -2, -4, 9].
"+" - Add -4 + 9 = 5 to the record, record is now [5, -2, -4, 9, 5].
"+" - Add 9 + 5 = 14 to the record, record is now [5, -2, -4, 9, 5, 14].
The total sum is 5 + -2 + -4 + 9 + 5 + 14 = 27.
Example 3:

Input: ops = ["1","C"]
Output: 0
Explanation:
"1" - Add 1 to the record, record is now [1].
"C" - Invalidate and remove the previous score, record is now [].
Since the record is empty, the total sum is 0.


Constraints:

1 <= operations.length <= 1000
operations[i] is "C", "D", "+", or a string representing an integer in the range [-3 * 104, 3 * 104].
For operation "+", there will always be at least two previous scores on the record.
For operations "C" and "D", there will always be at least one previous score on the record.
*/

func sum(arr []int) int {
	sum := 0
	for _, ele := range arr {
		sum = sum + ele
	}
	return sum
}

func sumLastTwo(arr []int) int {
	sum := 0
	if len(arr) == 0 {
		return sum
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		sum = arr[len(arr)-1] + arr[len(arr)-2]
		return sum
	}
}

// TC O(N), SC - O(N)
func calPoints(arr []string) int {
	N := len(arr)
	stack := []int{}
	for i := 0; i < N; i++ {
		if arr[i] == "+" {
			sum := sumLastTwo(stack)
			stack = append(stack, sum)

		} else if arr[i] == "D" {
			if len(stack) != 0 {
				val := stack[len(stack)-1]
				stack = append(stack, 2*val)
			}
		} else if arr[i] == "C" {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			num, _ := strconv.Atoi(arr[i])
			stack = append(stack, num)
		}

	}
	return sum(stack)
}

func mainBaseball() {
	arr := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}

	fmt.Println(calPoints(arr))
}
