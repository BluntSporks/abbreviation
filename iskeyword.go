package abbr

// IsKeyword checks if a word is a keyword of the given language.
// Note: Should lowercase words before passing them to this function.
func IsKeyword(lang string, word string) bool {
	if Keywords[lang][word] {
		return true
	}
	return false
}
