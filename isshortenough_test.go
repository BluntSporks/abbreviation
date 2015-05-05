package abbr

import "testing"

func TestIsShortEnough(t *testing.T) {
	if IsShortEnough("cat", "catty") {
		t.Error("Expected false, got true")
	}
	if !IsShortEnough("cat", "cation") {
		t.Error("Expected true, got false")
	}
}
