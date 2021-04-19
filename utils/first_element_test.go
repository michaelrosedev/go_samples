package utils

import (
	"fmt"
	"testing"
)

func TestWhenThereIsOnlyOneUniqueElement(t *testing.T) {
	input := []string{
		"elem1",
		"elem1",
		"elem2",
		"elem3",
		"elem3",
	}

	expected := "elem2"

	result := FirstUniqueElement(input)

	testResult(t, result, expected)
}

func TestWhenThereAreTwoUniqueElements(t *testing.T) {
	input := []string{
		"elem1",
		"elem2",
		"elem2",
		"elem3",
		"elem4",
		"elem4",
	}

	expected := "elem1"

	result := FirstUniqueElement(input)

	testResult(t, result, expected)
}

func TestWhenThereAreNoUniqueElements(t *testing.T) {
	input := []string{
		"elem1",
		"elem1",
		"elem2",
		"elem2",
		"elem3",
		"elem3",
	}

	result := FirstUniqueElement(input)

	testResult(t, result, "")
}

func TestWhenInputIsNil(t *testing.T) {
	result := FirstUniqueElement(nil)

	testResult(t, result, "")
}

func TestWhenInputHasNoElement(t *testing.T) {
	input := []string{}

	result := FirstUniqueElement(input)

	testResult(t, result, "")
}

func testResult(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected unique element of %q but got %q", expected, result)
	}
}

func ExampleFirstUniqueElement_single() {
	input := []string{
		"chicken",
		"chicken",
		"chicken",
		"horse", //unique
		"chicken",
	}
	result := FirstUniqueElement(input)
	fmt.Print(result)
	// Output: horse
}

func ExampleFirstUniqueElement_multiple() {
	input := []string{
		"apple", //unique
		"banana",
		"banana",
		"carrot",
		"carrot",
		"dartboard",
		"dartboard",
		"egg", //unique
		"fish",
		"fish",
	}
	result := FirstUniqueElement(input)
	fmt.Print(result)
	// Output: apple
}
