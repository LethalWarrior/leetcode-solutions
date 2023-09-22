package romantointeger

func RomanToInt(s string) int {
	romanNumerals := [256]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := romanNumerals[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if romanNumerals[s[i]] < romanNumerals[s[i+1]] {
			result -= romanNumerals[s[i]]
			continue
		}
		result += romanNumerals[s[i]]
	}

	return result
}
