package abbr

// IsLongEnough checks if the short form is at least three characters long.
func IsLongEnough(short string) bool {
	return len(short) >= 3
}
