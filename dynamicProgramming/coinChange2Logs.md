i 0 acc 0
i 0 acc 1
i 0 acc 2
i 0 acc 3
i 0 acc 4
i 0 acc 5
return 1
i 1 acc 4
i 1 acc 6
return 0
i 2 acc 4
i 2 acc 9
return 0
i 3 acc 4
return 0
return cache[key] [2 4] 0
return cache[key] [1 4] 0
return cache[key] [0 4] 1
i 1 acc 3
i 1 acc 5
return 1
i 2 acc 3
i 2 acc 8
return 0
i 3 acc 3
return 0
return cache[key] [2 3] 0
return cache[key] [1 3] 1
return cache[key] [0 3] 2
i 1 acc 2
i 1 acc 4
return val cache[key] [1 4] 0
i 2 acc 2
i 2 acc 7
return 0
i 3 acc 2
return 0
return cache[key] [2 2] 0
return cache[key] [1 2] 0
return cache[key] [0 2] 2
i 1 acc 1
i 1 acc 3
return val cache[key] [1 3] 1
i 2 acc 1
i 2 acc 6
return 0
i 3 acc 1
return 0
return cache[key] [2 1] 0
return cache[key] [1 1] 1
return cache[key] [0 1] 3
i 1 acc 0
i 1 acc 2
return val cache[key] [1 2] 0
i 2 acc 0
i 2 acc 5
return 1
i 3 acc 0
return 0
return cache[key] [2 0] 1
return cache[key] [1 0] 1
return cache[key] [0 0] 4
Number of ways to make change: 4

```go
func dfs(i, acc, amount int, coins []int, cache map[[2]int]int) int {
	fmt.Println("i", i, "acc", acc)
	if acc == amount {
		fmt.Println("return 1")
		return 1
	}
	if acc > amount || i == len(coins) {
		fmt.Println("return 0")
		return 0
	}
	key := [2]int{i, acc}
	if val, exists := cache[key]; exists {
		fmt.Println("return val cache[key]", key, cache[key])
		return val
	}

	cache[key] = dfs(i, acc+coins[i], amount, coins, cache) + dfs(i+1, acc, amount, coins, cache)
	fmt.Println("return cache[key]", key, cache[key])
	return cache[key]
}
```

```go
        COINS
        1 2 5 0
A [0]: [1,1,1,1]
M [1]: [1,0,0,0]
O [2]: [2,1,0,0]
U [3]: [2,0,0,0]
N [4]: [3,1,0,0]
T [5]: [4,1,1,0]
```
