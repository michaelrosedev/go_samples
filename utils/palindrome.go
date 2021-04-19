package utils

import "strings"

// This function will determine if the given string is a palindrome or not.
//
// A palindrome is a word that is spelled the same backwards as it is forwards, e.g. "racecar"
func IsPalindrome(input string) bool {

	li := strings.ToLower(input)

	chars := len(input)

	if chars == 0 {
		return false
	}

	for i := 0; i < chars/2; i++ {
		start := li[i]
		end := li[(chars-1)-i]

		if start != end {
			return false
		}
	}

	return true

}
