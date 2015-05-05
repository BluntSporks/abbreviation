package abbr

// HasVowel checks if the short form has a vowel.
func HasVowel(short string) bool {
	for _, char := range short {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'y':
			return true
		}

	}
	return false
}
