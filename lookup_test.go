package abbr

import "testing"

func TestLookUp(t *testing.T) {
	// Test successes.
	exp1 := "abbrev"
	act1 := LookUp("abbreviate")
	if exp1 != act1 {
		t.Error("Expected", exp1, "got", act1)
	}
	act2 := LookUp("abbreviates")
	if exp1 != act2 {
		t.Error("Expected", exp1, "got", act2)
	}
	act3 := LookUp("abbreviated")
	if exp1 != act3 {
		t.Error("Expected", exp1, "got", act3)
	}
	act4 := LookUp("abbreviating")
	if exp1 != act4 {
		t.Error("Expected", exp1, "got", act4)
	}
	act5 := LookUp("abbreviation")
	if exp1 != act5 {
		t.Error("Expected", exp1, "got", act5)
	}

	// Test fails.
	exp2 := ""
	act6 := LookUp("abbrevia")
	if exp2 != act6 {
		t.Error("Expected", exp2, "got", act6)
	}
	act7 := LookUp("xyz")
	if exp2 != act7 {
		t.Error("Expected", exp2, "got", act7)
	}
}
