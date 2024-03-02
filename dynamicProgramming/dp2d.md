# helperM

cols R, C 0 0
cols R, C 0 1
cols R, C 0 2
rows R, C 0 3
rows R, C 1 3
rows R, C 2 3
rows R, C 2 3
paths 1
rows R, C 1 3
paths 1
rows R, C 0 3
paths 1
cols R, C 0 2
paths 1
rows R, C 0 2
cols R, C 1 2
cols R, C 1 2
paths 1
rows R, C 1 2
cols R, C 2 2
cols R, C 2 2
paths 1
rows R, C 2 2
cols R, C 3 2
cols R, C 3 2
paths 1
rows R, C 2 2
paths 2
rows R, C 1 2
paths 3
rows R, C 0 2
paths 4
cols R, C 0 1
paths 4
rows R, C 0 1
cols R, C 1 1
cols R, C 1 1
paths 3
rows R, C 1 1
cols R, C 2 1
cols R, C 2 1
paths 2
rows R, C 2 1
cols R, C 3 1
cols R, C 3 1
paths 1
rows R, C 2 1
paths 3
rows R, C 1 1
paths 6
rows R, C 0 1
paths 10
cols R, C 0 0
paths 10
rows R, C 0 0
cols R, C 1 0
cols R, C 1 0
paths 6
rows R, C 1 0
cols R, C 2 0
cols R, C 2 0
paths 3
rows R, C 2 0
cols R, C 3 0
cols R, C 3 0
paths 1
rows R, C 2 0
paths 4
rows R, C 1 0
paths 10
rows R, C 0 0
paths 20

```go
func helperM(R, C int, cache *[][]int, ROWS, COLS int) int {
	// Base case: Destination reached
	if R == ROWS-1 && C == COLS-1 {
		return 1
	}

	// If already calculated, return the cached value
	if (*cache)[R][C] != 0 {
		return (*cache)[R][C]
	}

	// Initialize paths count
	paths := 0

	// Explore rightward move if within bounds
	if C < COLS-1 {
		fmt.Println("cols R, C", R, C)
		paths += helperM(R, C+1, cache, ROWS, COLS)
		fmt.Println("cols R, C", R, C)
		fmt.Println("paths", paths)
	}

	// Explore downward move if within bounds
	if R < ROWS-1 {
		fmt.Println("rows R, C", R, C)
		paths += helperM(R+1, C, cache, ROWS, COLS)
		fmt.Println("rows R, C", R, C)
		fmt.Println("paths", paths)
	}

	// Update cache with the sum of paths from downward and rightward moves
	(*cache)[R][C] = paths

	return paths
}
```
