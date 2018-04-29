package main

import (
	"testing"
)

func TestReverseToReturnReversedInputString(t *testing.T) {
	actual := Reverse("abcd")
	expected := "dcba"

	if Reverse("abcd") != expected {
		t.Errorf("Expected output be %v, but got %v", expected, actual)
	}

	actual = Reverse("Roman")
	expected = "namoR"

	if actual != "namoR" {
		t.Errorf("Expected output be %v, but got %v", expected, actual)
	}
}
