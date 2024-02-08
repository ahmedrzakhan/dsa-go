package main

import (
	"fmt"
)

/**
752. Open the Lock

You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3',
'4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn
'9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels
 of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total
number of turns required to open the lock, or -1 if it is impossible.

Example 1:

Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".
Example 2:

Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".
Example 3:

Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation: We cannot reach the target without getting stuck.

Constraints:

1 <= deadends.length <= 500
deadends[i].length == 4
target.length == 4
target will not be in the list deadends.
target and deadends[i] consist of digits only.
*/

// TC - O(N*M)/O(1), SC - O(N*M)/O(1), N(1000),M (8) is combinations
func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	if dead["0000"] {
		return -1
	}

	queue := []string{"0000"}
	visited := make(map[string]bool)
	visited["0000"] = true

	var count int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == target {
				return count
			}

			for pos := 0; pos < 4; pos++ {
				for _, move := range []int{-1, 1} {
					next := turn(curr, pos, move)
					if !visited[next] && !dead[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		count++
	}

	return -1
}

func turn(s string, pos, move int) string {
	digits := []rune(s)
	digits[pos] = rune((int(digits[pos])-'0'+move+10)%10 + '0')
	return string(digits)
}

func mainOL() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target)) // Expected output: 6
}
