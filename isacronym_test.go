package abbr

import "testing"

func TestIsAcronym(t *testing.T) {
	if !IsAcronym("tla", "three letter acronym") {
		t.Error("Expected true, got false")
	}
	if !IsAcronym("tla", "three  letter  acronym") {
		t.Error("Expected true, got false")
	}
	if IsAcronym("tlal", "three letter acronym") {
		t.Error("Expected false, got true")
	}
	if IsAcronym("tl", "three letter acronym") {
		t.Error("Expected false, got true")
	}
	if IsAcronym("tla", "three acronym letters") {
		t.Error("Expected false, got true")
	}

}
