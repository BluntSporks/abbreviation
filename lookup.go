package abbr

import (
	"strings"
)

// LookUp looks up a word to find its short form, if it is known.
// Note: Should lowercase word before passing them to this function.
func LookUp(word string) string {
	for long, short := range Abbrevs {
		if strings.Contains(word, long) {
			return short
		}
	}
	return ""
}
