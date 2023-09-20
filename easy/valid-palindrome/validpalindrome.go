package validpalindrome

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[\s\W_]`)
	s = re.ReplaceAllString(s, "")

	j := len(s) - 1

	for i := 0; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}

	return true
}
