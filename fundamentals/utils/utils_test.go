package utils

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		word       string
		palindrome bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			result := CheckPalindrome(tt.word)
			if result != tt.palindrome {
				t.Errorf("CheckPalindrome(%q) = %v; want %v", tt.word, result, tt.palindrome)
			}
		})
	}
}

func CompareMap(map1 map[string]int, map2 map[string]int) bool {
	for k, v := range map1 {
		if v != map2[k] {
			return false
		}
	}
	for k, v := range map2 {
		if v != map1[k] {
			return false
		}
	}
	return true
}

func TestCount(t *testing.T) {

	tests := []struct {
		word  string
		count map[string]int
	}{
		{"rr", map[string]int{"r": 2}},
		{"hello", map[string]int{"h": 1, "e": 1, "l": 2, "o": 1}},
		{"A m", map[string]int{"A": 1, "m": 1}},
		{"ra!-", map[string]int{"r": 1, "a": 1}},
	}

	for _, tst := range tests {
		t.Run(tst.word, func(t *testing.T) {
			result := CountFrequency(tst.word)
			if !CompareMap(result, tst.count) {
				t.Errorf("Count(%s)", tst.word)
			}
		})
	}
}
