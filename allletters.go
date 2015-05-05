package abbr

// AllLetters checks that all of the letters in the short form are present in the long form, and in the same order.
func AllLetters(short string, long string) bool {
	chars := []rune(long)
	i := 0
	for _, char := range short {
		for i < len(long) && char != chars[i] {
			i++
		}
		if i == len(long) {
			return false
		}
	}
	return true
}
