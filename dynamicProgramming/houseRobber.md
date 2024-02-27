i before rob 0
i before rob 2
i 2
robCurrent 3
i before skip 2
i 2
skipCurrent 0
max(robCurrent, skipCurrent) 3
i 0
robCurrent 4
i before skip 0
i before rob 1
i 1
robCurrent 2
i before skip 1
i before rob 2
i 2
robCurrent 3
i before skip 2
i 2
skipCurrent 0
max(robCurrent, skipCurrent) 3
i 1
skipCurrent 3
max(robCurrent, skipCurrent) 3
i 0
skipCurrent 3
max(robCurrent, skipCurrent) 4
Brute Force: 4

```go
func helper(A []int, i int) int {
	if i >= len(A) {
		return 0
	}

	// Option 1: Rob the current house
	fmt.Println("i before rob", i)
	robCurrent := A[i] + helper(A, i+2) // Add loot and skip the next
	fmt.Println("i", i)
	fmt.Println("robCurrent", robCurrent)

	// Option 2: Skip the current house
	fmt.Println("i before skip", i)
	skipCurrent := helper(A, i+1) // Move to the next house
	fmt.Println("i", i)
	fmt.Println("skipCurrent", skipCurrent)

	fmt.Println("max(robCurrent, skipCurrent)", max(robCurrent, skipCurrent))
	return max(robCurrent, skipCurrent)
}
```
