index : i 0 0
firstPartSum 7
i+1, k-1 1 2
index : i 1 1
firstPartSum 2
i+1, k-1 2 1
largestSum 23
minLargestSum 23
index : i 1 2
firstPartSum 7
i+1, k-1 3 1
largestSum 18
minLargestSum 18
index : i 1 3
firstPartSum 17
i+1, k-1 4 1
largestSum 17
minLargestSum 17
largestSum 17
minLargestSum 17
index : i 0 1
firstPartSum 9
i+1, k-1 2 2
index : i 2 2
firstPartSum 5
i+1, k-1 3 1
largestSum 18
minLargestSum 18
index : i 2 3
firstPartSum 15
i+1, k-1 4 1
largestSum 15
minLargestSum 15
largestSum 15
minLargestSum 15
index : i 0 2
firstPartSum 14
i+1, k-1 3 2
index : i 3 3
firstPartSum 10
i+1, k-1 4 1
largestSum 10
minLargestSum 10
largestSum 14
minLargestSum 14
Minimum largest split sum: 14

```go
func split(arr []int, index, K int) int {
	if K == 1 {
		return sum(arr[index:])
	}
	minLargestSum := math.MaxInt
	for i := index; i <= len(arr)-K; i++ {
		fmt.Println("index : i", index, i)
		firstPartSum := sum(arr[index : i+1])
		fmt.Println("firstPartSum", firstPartSum)
		fmt.Println("i+1, k-1", i+1, K-1)
		largestSum := max(firstPartSum, split(arr, i+1, K-1))
		fmt.Println("largestSum", largestSum)
		minLargestSum = min(minLargestSum, largestSum)
		fmt.Println("minLargestSum", minLargestSum)
		if firstPartSum > minLargestSum {
			break
		}
	}
	return minLargestSum
}
```
