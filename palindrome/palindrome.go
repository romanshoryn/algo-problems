package main

import (
	"strings"
)

// Palindrome returns true if input string is palindrome
func Palindrome(s string) bool {
	str := strings.ToLower(s)
	str = strings.Replace(str, " ", "", -1)

	arr := []byte(str)
	mid := len(arr) / 2

	for i, j := 0, len(arr)-1; i < mid; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}
