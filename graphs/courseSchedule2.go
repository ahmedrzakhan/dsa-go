package main

import "fmt"

/**
210. Course Schedule II

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you
must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid
answers, return any of them. If it is impossible to finish all courses, return an empty array.

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished
course 0. So the correct course order is [0,1].
Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished
both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
Example 3:

Input: numCourses = 1, prerequisites = []
Output: [0]

Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
All the pairs [ai, bi] are distinct.
*/

// TC - O(V+E), SC - O(V+E)
func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)
	result := make([]int, 0)
	visited := make(map[int]bool, numCourses)  // Tracks visited courses
	visiting := make(map[int]bool, numCourses) // Tracks courses currently being visited (for cycle detection)

	// Build adjacency list
	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			if !dfsCS2(i, visiting, visited, adjList, &result) {
				return []int{} // Cycle detected or cannot finish all courses
			}
		}
	}

	// Reverse the result to get the correct order
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - 1 - i // Calculate the counterpart from the end
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func dfsCS2(course int, visiting, visited map[int]bool, adjList map[int][]int, result *[]int) bool {
	if visiting[course] { // Cycle detected
		return false
	}
	if visited[course] { // Already processed
		return true
	}

	visiting[course] = true
	for _, pre := range adjList[course] {
		if !dfsCS2(pre, visiting, visited, adjList, result) {
			return false // Cycle detected in deeper recursion
		}
	}

	visiting[course] = false          // Backtrack
	visited[course] = true            // Mark as processed
	*result = append(*result, course) // Add course to result

	return true
}

func mainCS2() {
	numOfCourses5 := 5
	prerequisites5 := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4}}
	fmt.Println(findOrder(numOfCourses5, prerequisites5)) // Expected: Any valid topological order, e.g., [2, 4, 3, 1, 0]

	numOfCourses3 := 3
	prerequisites3 := [][]int{{1, 0}, {2, 1}, {0, 2}}
	fmt.Println(findOrder(numOfCourses3, prerequisites3)) // Expected: [] since there's a cycle
}
