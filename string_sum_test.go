package string_sum

import (
	"testing"
)

func TestStringSum1(t *testing.T) {
	want := "-7"
	got, err := StringSum("      - 15 + 8  ")

	if got != want {
		t.Errorf("TestStringSum = %s, want %s\nerror: %v", got, want, err)
	}
}

func TestStringSum2(t *testing.T) {
	want := "23"
	got, err := StringSum("       15 + 8  ")

	if got != want {
		t.Errorf("TestStringSum2 = %s, want %s\nerror: %v", got, want, err)
	}
}
