                          [0-5, Sum: 36]
                         /             \
              [0-2, Sum: 9]             [3-5, Sum: 27]
              /       \                   /        \
       [0-1, Sum: 4]  [2, Sum: 5]   [3-4, Sum: 16]  [5, Sum: 11]
        /     \                        /    \
[0, Sum: 1] [1, Sum: 3]        [3, Sum: 7] [4, Sum: 9]


// logs
left Build, L: 0 M: 2 R: 5
left Build, L: 0 M: 1 R: 2
left Build, L: 0 M: 0 R: 1
root.left.sum 1
right Build, L: 0 M+1: 1 R: 1
root.right.sum 3
root.sum 4
root.left.sum 4
right Build, L: 0 M+1: 2 R: 2
root.right.sum 5
root.sum 9
root.left.sum 9
right Build, L: 0 M+1: 3 R: 5
left Build, L: 3 M: 4 R: 5
left Build, L: 3 M: 3 R: 4
root.left.sum 7
right Build, L: 3 M+1: 4 R: 4
root.right.sum 9
root.sum 16
root.left.sum 16
right Build, L: 3 M+1: 5 R: 5
root.right.sum 11
root.sum 27
root.right.sum 27
root.sum 36
Initial sum in range [1, 3]: 15
Updated sum in range [1, 3]: 22

func Build(nums []int, L, R int) *SegmentTree {
	if L == R {
		return &SegmentTree{sum: nums[L], L: L, R: R}
	}

	M := L + (R-L)/2
	root := &SegmentTree{L: L, R: R}
	fmt.Println("left Build, L:", L, "M:", M, "R:", R)
	root.left = Build(nums, L, M)
	fmt.Println("root.left.sum", root.left.sum)
	fmt.Println("right Build, L:", L, "M+1:", M+1, "R:", R)
	root.right = Build(nums, M+1, R)
	fmt.Println("root.right.sum", root.right.sum)

	root.sum = root.left.sum + root.right.sum
	fmt.Println("root.sum", root.sum)
	return root
}