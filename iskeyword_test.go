package abbr

import "testing"

func TestIsKeyword(t *testing.T) {
	if !IsKeyword("go", "goto") {
		t.Error("Expected true, got false")
	}
	if !IsKeyword("php", "goto") {
		t.Error("Expected true, got false")
	}
	if IsKeyword("java", "goto") {
		t.Error("Expected false, got true")
	}
}
