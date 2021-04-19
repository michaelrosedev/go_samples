package utils

import "fmt"

// The classic. This function takes an input number and:
//
// Will return "fizz" if the number is divisible by 3
//
// Will return "buzz" if the number is divisible by 5
//
// Will return "fizzbuzz" if the number is divisible by both 3 and 5
//
// Will return the input as a string otherwise
func FizzBuzz(input int) string {

	result := ""

	if input%3 == 0 {
		result += "fizz"
	}

	if input%5 == 0 {
		result += "buzz"
	}

	if result == "" {
		result = fmt.Sprint(input)
	}

	return result
}
