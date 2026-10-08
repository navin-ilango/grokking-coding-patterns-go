package main

import "unicode"

// Examples
// Input: "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

func isPalindrome(s string) bool {

	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNumeric(rune(s[left])) {
			left++
		}
		for left < right && !isAlphaNumeric(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
