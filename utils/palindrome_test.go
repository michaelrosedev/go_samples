package utils

import (
	"fmt"
	"testing"
)

func TestWhenIsAPalindromeOfAllSameCharacters(t *testing.T) {
	input := "aaaaa"
	expected := true

	actual := IsPalindrome(input)

	if actual != expected {
		t.Errorf("Expected %q to be identified as a palindrome, but it wasn't", input)
	}
}

func TestWhenNotAPalindrome(t *testing.T) {
	input := "12345"
	expected := false

	actual := IsPalindrome(input)

	if actual != expected {
		t.Errorf("Expected %q to not be identified as a palindrome, but it was", input)
	}
}

func TestWhenValidPalindrome(t *testing.T) {
	input := "cattac"
	expected := true

	actual := IsPalindrome(input)

	if actual != expected {
		t.Errorf("Expected %q to be identified as a palindrome, but it wasn't", input)
	}
}

func TestWhenInputIsEmpty(t *testing.T) {
	input := ""
	expected := false

	actual := IsPalindrome(input)

	if actual != expected {
		t.Errorf("Expected '%q' to not be identified as a palindrome, but it was", input)
	}
}

func TestWhenInputIsMixedCase(t *testing.T) {
	input := "CatBigGibTac"
	expected := true

	actual := IsPalindrome(input)

	if actual != expected {
		t.Errorf("Expected '%q' to be identified as a palindrome, but it wasn't", input)
	}
}

func ExampleIsPalindrome_true() {
	input := "racecar"

	result := IsPalindrome(input)
	fmt.Print(result)
	// Output: true
}

func ExampleIsPalindrome_false() {
	input := "fishing"

	result := IsPalindrome(input)
	fmt.Print(result)
	// Output: false
}
