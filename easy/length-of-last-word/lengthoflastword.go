package lengthoflastword

func LengthOfLastWord(s string) int {
	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i -= 1
	}

	lastIdx := i

	for i >= 0 && s[i] != ' ' {
		i -= 1
	}

	return lastIdx - i
}
