package abbr

import "testing"

func TestIsLongEnough(t *testing.T) {
	if IsLongEnough("do") {
		t.Error("Expected false, got true")
	}
	if !IsLongEnough("dog") {
		t.Error("Expected true, got false")
	}
}
