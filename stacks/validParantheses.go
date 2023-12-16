package main

import "fmt"

/**
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

*/

/**
If number of items odd in total return unbalanced immediately at the start

Ask interviewer if anything other than these six characters
*/

func includes(arr []string, val string) bool {
	N := len(arr)
	for i := 0; i < N; i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

// TC - O(N), SC - (O(N))
func isValid(s string) bool {
	validSym := []string{"(", "{", "["}
	stack := []string{}

	for i := 0; i < len(s); i++ {
		val := string(s[i])
		if includes(validSym, val) {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				lastElem := stack[len(stack)-1]
				if val == "]" && lastElem != "[" || val == "}" && lastElem != "{" ||
					val == ")" && lastElem != "(" {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}[]"))
}
