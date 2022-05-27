package string_sum

import (
	"testing"
)

func TestStringSumSubstractOperation(t *testing.T) {
	want := "7"
	got, err := StringSum("15-8")

	if got != want {
		t.Errorf("got: %s, want: %s\nerror: %v", got, want, err)
	}
}

func TestStringSumAddOperation(t *testing.T) {
	want := "23"
	got, err := StringSum("15+8")

	if got != want {
		t.Errorf("got: %s, want: %s\nerror: %v", got, want, err)
	}
}

func TestStringSumErrorNotTwoOperands(t *testing.T) {
	want := errorNotTwoOperands
	_, err := StringSum("15")

	if err.Error() != want.Error() {
		t.Errorf("got: %s, want: %s", err, want)
	}
}

func TestStringSumMoreThanTwoOperands(t *testing.T) {
	want := errorNotTwoOperands
	_, err := StringSum("15+8+9")

	if err.Error() != want.Error() {
		t.Errorf("got: %s, want: %s", err, want)
	}
}

func TestStringSumErrorEmptyInput(t *testing.T) {
	want := errorEmptyInput
	_, err := StringSum("          ")

	if err.Error() != want.Error() {
		t.Errorf("got: %s, want: %s", err, want)
	}
}
