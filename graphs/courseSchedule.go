package main

import "fmt"

/**
207. Course Schedule

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an
array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you
 want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished
course 1. So it is impossible.

Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
All the pairs prerequisites[i] are unique.
*/

// TC - O(V+E), SC - O(V+E)
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	visited := make(map[int]struct{})

	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	for i := 0; i < numCourses; i++ {
		if !dfsCS(i, graph, visited) {
			return false
		}
	}

	return true
}

func dfsCS(course int, graph map[int][]int, visited map[int]struct{}) bool {
	if _, ok := visited[course]; ok {
		return false
	}

	if len(graph[course]) == 0 {
		return true
	}

	visited[course] = struct{}{}

	for _, pre := range graph[course] {
		if !dfsCS(pre, graph, visited) {
			return false
		}
	}
	delete(visited, course)
	graph[course] = []int{}

	return true
}

func mainCS() {
	// numCourses := 2
	// prerequisites := [][]int{{1, 0}}
	// fmt.Println(canFinish(numCourses, prerequisites)) // True if all courses can be finished, False otherwise
	// prerequisitesF := [][]int{{1, 0}, {0, 1}}
	// numCoursesF := 2
	// fmt.Println(canFinish(numCoursesF, prerequisitesF)) // True if all courses can be finished, False otherwise

	numOfCourses5 := 5
	prerequisites5 := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4}}
	fmt.Println(canFinish(numOfCourses5, prerequisites5))

	numOfCourses3 := 3
	prerequisites3 := [][]int{{1, 0}, {2, 1}, {0, 2}}
	fmt.Println(canFinish(numOfCourses3, prerequisites3))
}
