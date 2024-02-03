package main

import "fmt"

type SegmentTree struct {
	sum   int
	left  *SegmentTree
	right *SegmentTree
	L, R  int
}

// TC - O(N)
// Build function constructs the Segment Tree recursively
func Build(nums []int, L, R int) *SegmentTree {
	if L == R {
		return &SegmentTree{sum: nums[L], L: L, R: R}
	}

	M := L + (R-L)/2
	root := &SegmentTree{L: L, R: R}
	root.left = Build(nums, L, M)
	root.right = Build(nums, M+1, R)
	root.sum = root.left.sum + root.right.sum
	return root
}

// TC - O(LogN)
// Update function updates the value at a specific index
func (st *SegmentTree) Update(index, val int) {
	if st.L == st.R {
		st.sum = val
		return
	}

	M := st.L + (st.R-st.L)/2
	if index > M {
		st.right.Update(index, val)
	} else {
		st.left.Update(index, val)
	}
	st.sum = st.left.sum + st.right.sum
}

// TC - O(LogN)
// RangeQuery function performs a query for the sum in a range [L, R]
func (st *SegmentTree) RangeQuery(L, R int) int {
	if L == st.L && R == st.R {
		return st.sum
	}

	M := st.L + (st.R-st.L)/2
	if L > M {
		return st.right.RangeQuery(L, R)
	} else if R <= M {
		return st.left.RangeQuery(L, R)
	} else {
		return st.left.RangeQuery(L, M) + st.right.RangeQuery(M+1, R)
	}
}

func mainST() {
	nums := []int{1, 3, 5, 7, 9, 11}
	segmentTree := Build(nums, 0, len(nums)-1)

	fmt.Println("Initial sum in range [1, 3]:", segmentTree.RangeQuery(1, 3))

	// Update value at index 1 to 10
	segmentTree.Update(1, 10)
	fmt.Println("Updated sum in range [1, 3]:", segmentTree.RangeQuery(1, 3))
}
