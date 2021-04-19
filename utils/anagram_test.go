package utils

import (
	"fmt"
	"testing"
)

func TestWhenIsAnagram(t *testing.T) {
	input1 := "sadder"
	input2 := "dreads"
	expected := true

	result := IsAnagram(input1, input2)

	assertExpectTrue(t, result, expected, input1, input2)

}

func TestWhenNotAnagram(t *testing.T) {
	input1 := "potato"
	input2 := "badger"
	expected := false

	result := IsAnagram(input1, input2)

	assertExpectFalse(t, result, expected, input1, input2)
}

func TestWhenAlmostAnAnagram(t *testing.T) {
	input1 := "carrot"
	input2 := "carrob"
	expected := false

	result := IsAnagram(input1, input2)

	assertExpectFalse(t, result, expected, input1, input2)
}

func TestWhenDifferentNumberOfLetters(t *testing.T) {
	input1 := "book"
	input2 := "character"
	expected := false

	result := IsAnagram(input1, input2)

	assertExpectFalse(t, result, expected, input1, input2)
}

func TestWhenSameCharactersButDifferentNumbers(t *testing.T) {
	input1 := "bookleg"
	input2 := "beeklog"
	expected := false

	result := IsAnagram(input1, input2)

	assertExpectFalse(t, result, expected, input1, input2)

}

func TestWhenAnagramWithSpaces(t *testing.T) {
	input1 := "a gentleman"
	input2 := "elegant man"
	expected := true

	result := IsAnagram(input1, input2)

	assertExpectTrue(t, result, expected, input1, input2)

}

func TestWhenAnagramWithOnly1InputHavingSpaces(t *testing.T) {
	input1 := "conversation"
	input2 := "voices rant on"
	expected := true

	result := IsAnagram(input1, input2)

	assertExpectTrue(t, result, expected, input1, input2)

}

func TestWhenAnagramButMixedCase(t *testing.T) {
	input1 := "TheMorseCODE"
	input2 := "HereComeDOTS"
	expected := true

	result := IsAnagram(input1, input2)

	assertExpectTrue(t, result, expected, input1, input2)

}

func assertExpectTrue(t testing.TB, result, expected bool, input1, input2 string) {
	if result != expected {
		t.Errorf("Expected %q to be an anagram of %q", input1, input2)
	}
}

func assertExpectFalse(t testing.TB, result, expected bool, input1, input2 string) {
	if result != expected {
		t.Errorf("Did not expect %q to be an anagram of %q", input1, input2)
	}
}

func ExampleIsAnagram() {
	input1 := "conversation"
	input2 := "voices rant on"

	result := IsAnagram(input1, input2)
	fmt.Println(result)
	// Ouput: true
}
