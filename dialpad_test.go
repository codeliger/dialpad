package main

import (
	"reflect"
	"testing"
)

func TestNumToStrings(t *testing.T) {
	expected := []string{
		"ADG", "ADH", "ADI", "AEG", "AEH", "AEI", "AFG", "AFH", "AFI",
		"BDG", "BDH", "BDI", "BEG", "BEH", "BEI", "BFG", "BFH", "BFI",
		"CDG", "CDH", "CDI", "CEG", "CEH", "CEI", "CFG", "CFH", "CFI",
	}

	actual := num2strings("234", lookup)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpect %v\nActual %v", expected, actual)
	}
}
