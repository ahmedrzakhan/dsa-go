package main

import (
	"fmt"
)

func mainHash() {
	names := []string{"alice", "brad", "collin", "brad", "dylan", "kim"}
	countMap := make(map[string]int)

	for _, name := range names {
		countMap[name]++
	}
	fmt.Println(countMap)
}
