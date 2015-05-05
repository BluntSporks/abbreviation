package abbr

import "testing"

func TestHasVowel(t *testing.T) {
	if !HasVowel("cat") {
		t.Error("Expected true, got false")
	}
	if !HasVowel("egg") {
		t.Error("Expected true, got false")
	}
	if !HasVowel("go") {
		t.Error("Expected true, got false")
	}
	if !HasVowel("i") {
		t.Error("Expected true, got false")
	}
	if !HasVowel("why") {
		t.Error("Expected true, got false")
	}
	if HasVowel("bcdfghjklmnpqrstvwxz") {
		t.Error("Expected false, got true")
	}

}
