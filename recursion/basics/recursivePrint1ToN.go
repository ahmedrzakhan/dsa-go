package main

import (
	"fmt"
)

// recursivePrint prints numbers from 1 to n.
func recursivePrint(n int) {
	if n <= 0 {
		return
	}
	recursivePrint(n - 1)
	fmt.Println(n)
}

func mainRecPrin() {
	N := 7 // Change this to print up to any number
	recursivePrint(N)
}
