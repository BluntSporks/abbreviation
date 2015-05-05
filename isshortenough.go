package abbr

// IsShortEnough checks if the short form is half or less in length compared to the long form.
func IsShortEnough(short string, long string) bool {
	if len(short)*2 > len(long) {
		return false
	}
	return true
}
