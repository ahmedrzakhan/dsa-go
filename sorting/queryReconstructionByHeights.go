package main

import (
	"fmt"
	"sort"
)

/**
406. Queue Reconstruction by Height

You are given an array of people, people, which are the attributes of some people in a queue (not necessarily in order). Each people[i] = [hi, ki] represents the ith person of height hi with exactly ki other people in front who have a height greater than or equal to hi.

Reconstruct and return the queue that is represented by the input array people. The returned queue should be formatted as an array queue, where queue[j] = [hj, kj] is the attributes of the jth person in the queue (queue[0] is the person at the front of the queue).

Example 1:

Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
Explanation:
Person 0 has height 5 with no other people taller or the same height in front.
Person 1 has height 7 with no other people taller or the same height in front.
Person 2 has height 5 with two persons taller or the same height in front, which is person 0 and 1.
Person 3 has height 6 with one person taller or the same height in front, which is person 1.
Person 4 has height 4 with four people taller or the same height in front, which are people 0, 1, 2, and 3.
Person 5 has height 7 with one person taller or the same height in front, which is person 1.
Hence [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] is the reconstructed queue.
Example 2:

Input: people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
Output: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]

Constraints:

1 <= people.length <= 2000
0 <= hi <= 106
0 <= ki < people.length
It is guaranteed that the queue can be reconstructed.
*/

/**
Brute force: sort the people by their heights in descending order and, in case of a tie, by their k values in
ascending order. Then, we will insert each person into the queue at the index specified by their k value. This
approach works because, by inserting the tallest people first, every insertion operation is placed at the final
position of that person in the queue.
TC - O(N^2) coz of insertion after sorting
*/

// TC - O(N^2), SC - O(N)
func reconstructQueue(people [][]int) [][]int {
	// Sort the people to make the tallest person first, and in case of a tie, the one with fewer people in front comes first
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // Ascending order of k for same height
		}
		return people[i][0] > people[j][0] // Descending order of height
	})

	var queue [][]int
	for _, person := range people {
		// Insert each person at the index specified by their k value
		// Since the slice is 0-indexed, we insert directly at the k value
		index := person[1] // The k value specifies the insertion index
		queue = append(queue[:index], append([][]int{person}, queue[index:]...)...)
	}

	return queue
}

func mainQR() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructedQueue := reconstructQueue(people)
	fmt.Println("Reconstructed Queue:", reconstructedQueue)
	// Output: Reconstructed Queue: [[5 0] [7 0] [5 2] [6 1] [4 4] [7 1]]
}
