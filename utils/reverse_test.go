package utils

import (
	"fmt"
	"testing"
)

func TestWhenStringIsReversed(t *testing.T) {
	input := "abcdefg"
	expected := "gfedcba"

	result := Reverse(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestWhenStringIsAPalindrome(t *testing.T) {
	input := "bobbob"
	expected := "bobbob"

	result := Reverse(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestWhenReverseInputIsEmpty(t *testing.T) {
	input := ""
	expected := ""

	result := Reverse(input)

	if result != expected {
		t.Errorf("Expected an empty string but got %q", result)
	}
}

func ExampleReverse() {
	input := "abcdefgh"
	result := Reverse(input)
	fmt.Print(result)
	// Output: hgfedcba
}
