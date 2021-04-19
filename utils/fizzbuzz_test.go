package utils

import (
	"fmt"
	"testing"
)

func TestWhenInputOnlyDivisibleByThree(t *testing.T) {
	actual := FizzBuzz(3)

	expected := "fizz"

	if actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}

func TestWhenInputOnlyDivisibleByFive(t *testing.T) {
	actual := FizzBuzz(10)

	expected := "buzz"

	if actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}

func TestWhenInputDivisibleByBoth(t *testing.T) {
	actual := FizzBuzz(15)
	expected := "fizzbuzz"

	if actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}

func TestWhenInputDivisibleByNeither(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "2"

	if actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}

func TestOutputUpTo100(t *testing.T) {
	for i := 1; i <= 100; i++ {

		result := FizzBuzz(i)

		if i%3 == 0 && i%5 == 0 {
			if result != "fizzbuzz" {
				t.Errorf("Expected %q but got %q", "fizzbuzz", result)
			}
		}

		if i%3 == 0 && i%5 != 0 {
			if result != "fizz" {
				t.Errorf("Expected %q but got %q", "fizz", result)
			}
		}

		if i%3 != 0 && i%5 == 0 {
			if result != "buzz" {
				t.Errorf("Expected %q but got %q", "buzz", result)
			}
		}

		if i%3 != 0 && i%5 != 0 {
			if result != fmt.Sprint(i) {
				t.Errorf("Expected %q but got %q", fmt.Sprint(i), result)
			}
		}
	}
}

func ExampleFizzBuzz() {
	input := 15
	result := FizzBuzz(input)
	fmt.Print(result)
	// Output: fizzbuzz
}
