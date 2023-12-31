package main

import (
	"fmt"
)

// recursivePrint prints numbers from 1 to n.
func recursivePrintN(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	recursivePrintN(n - 1)
}

func mainRec() {
	N := 7 // Change this to print up to any number
	recursivePrintN(N)
}
