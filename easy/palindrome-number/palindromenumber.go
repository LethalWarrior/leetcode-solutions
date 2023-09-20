package easy

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strX := strconv.Itoa(x)

	j := len(strX) - 1

	for i := 0; i < j; i++ {
		if strX[i] != strX[j] {
			return false
		}
		j--
	}

	return true
}
