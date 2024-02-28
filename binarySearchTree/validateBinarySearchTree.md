```
   24
  /  \
 2   26
/ \
1  5
	\
	33
```

min -9223372036854775808
max 9223372036854775807
min -9223372036854775808
max 5
min 5
max 9223372036854775807
true
min -9223372036854775808
max 9223372036854775807
min -9223372036854775808
max 24
min -9223372036854775808
max 2
min 2
max 24
false

```go
func isValid(curr *TreeNode, min int, max int) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= min || curr.Val >= max {
		return false
	}
	fmt.Println("min", min)
	fmt.Println("max", max)
	return isValid(curr.Left, min, curr.Val) && isValid(curr.Right, curr.Val, max)
}
```
