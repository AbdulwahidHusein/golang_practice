package utils

import (
	"strings"
	"unicode"
)

// CountFrequency returns a map with the frequency of each letter in the given word.
func CountFrequency(word string) map[string]int {
	frequency := make(map[string]int)

	for _, letter := range word {
		if isAlphanumeric(letter) {
			frequency[string(letter)]++
		}
	}

	return frequency
}

func CheckPalindrome(word string) bool {
	left := 0
	right := len(word) - 1

	for left < right {
		for left < right && !isAlphanumeric(rune(word[left])) {
			left++
		}
		for left < right && !isAlphanumeric(rune(word[right])) {
			right--
		}
		if left < right && strings.ToLower(string(word[left])) != strings.ToLower(string(word[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
