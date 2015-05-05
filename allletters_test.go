package abbr

import "testing"

func TestAllLetters(t *testing.T) {
	if !AllLetters("a", "abc") {
		t.Error("Expected true, got false")
	}
	if !AllLetters("ace", "abcde") {
		t.Error("Expected true, got false")
	}
	if !AllLetters("abcde", "abcde") {
		t.Error("Expected true, got false")
	}
	if AllLetters("abcdef", "abcde") {
		t.Error("Expected false, got true")
	}
	if AllLetters("abdce", "abcde") {
		t.Error("Expected false, got true")
	}
	if AllLetters("abcdez", "abcde") {
		t.Error("Expected false, got true")
	}

}
