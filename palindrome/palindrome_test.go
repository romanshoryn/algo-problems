package main

import "testing"

func TestPalindromeToReturnTrueIfStrIsPalindrome(t *testing.T) {
	val := "abba"
	if Palindrome(val) != true {
		t.Errorf("Expected %v to be palindrom, but got false", val)
	}

	val = "Roman"
	if Palindrome(val) != false {
		t.Errorf("Expected %v to be palindrom, but got true", val)
	}

	val = "Air an aria"
	if Palindrome(val) != true {
		t.Errorf("Expected %v to be palindrom, but got false", val)
	}
}
