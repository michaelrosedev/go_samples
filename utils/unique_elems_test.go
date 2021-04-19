package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindUniqueElementWhenThereIsOne(t *testing.T) {
	input := []string{
		"elem1",
		"elem1",
		"elem1",
		"elem2",
		"elem2",
		"elem3",
		"elem4",
		"elem4",
		"elem4",
	}

	result := UniqueEntries(input)

	expected := []string{
		"elem3",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected to find element 'elem3' was unique, but didn't")
	}
}

func TestFindUniqueElementsWhenThereAreMultiple(t *testing.T) {
	input := []string{
		"elem1",
		"elem2",
		"elem3",
		"elem4",
	}

	result := UniqueEntries(input)

	expected := []string{
		"elem1",
		"elem2",
		"elem3",
		"elem4",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected all elements to be unique, but they weren't")
	}
}

func TestFindUniqueElementWhereThereAreNone(t *testing.T) {
	input := []string{
		"elem1",
		"elem1",
		"elem2",
		"elem2",
	}

	result := UniqueEntries(input)

	if result != nil {
		t.Error("Expected to find no unique elements")
	}
}

func ExampleUniqueEntries() {
	input := []string{
		"apple",
		"cat",
		"cat",
		"dart",
		"hat",
		"hat",
	}

	result := UniqueEntries(input)
	fmt.Println(result)
	// Output: [apple dart]
}
