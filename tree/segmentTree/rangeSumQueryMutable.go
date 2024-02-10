package main

import "fmt"

/**
307. Range Sum Query - Mutable

Given an integer array nums, handle multiple queries of the following types:

Update the value of an element in nums.
Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
void update(int index, int val) Updates the value of nums[index] to be val.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

Example 1:

Input
["NumArray", "sumRange", "update", "sumRange"]
[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
Output
[null, 9, null, 8]

Explanation
NumArray numArray = new NumArray([1, 3, 5]);
numArray.sumRange(0, 2); // return 1 + 3 + 5 = 9
numArray.update(1, 2);   // nums = [1, 2, 5]
numArray.sumRange(0, 2); // return 1 + 2 + 5 = 8

Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
0 <= index < nums.length
-100 <= val <= 100
0 <= left <= right < nums.length
At most 3 * 104 calls will be made to update and sumRange.
*/

type SegmentTreeRS struct {
	sum   int
	left  *SegmentTreeRS
	right *SegmentTreeRS
	L, R  int
}

type NumArray struct {
	root *SegmentTree
}

func BuildRS(nums []int, L, R int) *SegmentTree {
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

func ConstructorRS(nums []int) NumArray {
	return NumArray{root: Build(nums, 0, len(nums)-1)}
}

func (st *NumArray) UpdateRS(index int, val int) {
	st.root.UpdateRS(index, val)
}

func (st *SegmentTree) UpdateRS(index, val int) {
	if st.L == st.R {
		st.sum = val
		return
	}

	M := st.L + (st.R-st.L)/2
	if index <= M {
		st.left.Update(index, val)
	} else {
		st.right.Update(index, val)
	}
	st.sum = st.left.sum + st.right.sum
}

func (st *NumArray) SumRangeRS(left int, right int) int {
	return st.root.RangeQuery(left, right)
}

func (st *SegmentTreeRS) RangeQuery(L, R int) int {
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

func mainRS() {
	nums := []int{1, 3, 5, 7, 9, 11}
	numArray := ConstructorRS(nums)
	fmt.Println("Initial sum in range [1, 3]:", numArray.SumRangeRS(1, 3)) // Output: 15

	numArray.UpdateRS(1, 10)                                               // Update index 1 to value 10
	fmt.Println("Updated sum in range [1, 3]:", numArray.SumRangeRS(1, 3)) // Output: 22
}
