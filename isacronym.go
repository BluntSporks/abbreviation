package abbr

import "strings"

// IsAcronym checks if the short form is an acronym of the long.
func IsAcronym(short string, long string) bool {
	words := strings.Split(long, " ")
	j := 0
	for i := 0; i < len(short); i++ {
		// Skip any empty words.
		for j < len(words) && words[j] == "" {
			j++
		}

		// See if we are out of words.
		if j == len(words) {
			return false
		}

		// See if we have the wrong char for this word.
		if short[i] != words[j][0] {
			return false
		}
		j++

	}

	// See if we did not use up all the words.
	if j != len(words) {
		return false
	}

	// All tests passed.
	return true
}
