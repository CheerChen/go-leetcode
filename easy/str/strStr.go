package str

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	t := -1
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			t = i
			break
		}
	}
	return t
}
