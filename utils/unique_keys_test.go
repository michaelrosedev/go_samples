package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReturningUniqueKeys(t *testing.T) {
	input := map[string]int{
		"elem1": 1,
		"elem2": 10,
		"elem3": 1,
	}

	expected := map[string]int{
		"elem1": 1,
		"elem3": 1,
	}

	result := ToUniqueKeys(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to only contain unique keys from input")
	}
}

func TestWhenUniqueKeysInputIsNil(t *testing.T) {
	result := ToUniqueKeys(nil)

	if result != nil {
		t.Errorf("Expected result to be nil, but it wasn't")
	}
}

func ExampleToUniqueKeys() {
	input := map[string]int{
		"cats":    1,
		"dogs":    100,
		"fish":    1,
		"gerbils": 7,
	}
	result := ToUniqueKeys(input)
	fmt.Println(result)
	// Output: map[cats:1 fish:1]
}
