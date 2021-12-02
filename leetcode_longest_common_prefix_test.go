package leetcode_longest_common_prefix

import (
	"testing"
)

var testingStrings = [][]string {
	{"abc", "abcdefgh", "abcdefghijklmnop"},
	{"i love pasta", "i love pasta and wine"},
	{"gaaah", "waaaah", "aaah"},
	{"flower", "flow", "flight"},
	{"dog", "racecar", "car"},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, value := range testingStrings {
		retVal := longestCommonPrefix(value)
		if "" == retVal {
			t.Fatalf("Invalid answer. Given: %v. Recieved: %v.", value, retVal)
		}
	}
}