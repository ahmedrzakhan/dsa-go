package main

import "fmt"

// check if a given string is palindrome recursively

func checkPalindrome(str []rune, L, R int) bool {
	if str[L] != str[R] {
		return false
	}

	L++
	R--
	if L > R {
		return true
	}
	return checkPalindrome(str, L, R)
}

func mainPl() {
	str := "racecar"
	runes := []rune(str)
	L, R := 0, len(str)-1
	fmt.Println(checkPalindrome(runes, L, R))
}
