package indexoffirstoccurenceinstring

func StrStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)

	for i := 0; i < m-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
