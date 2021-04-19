package utils

// This function will reverse the characters in the provided input string
func Reverse(input string) string {
	var chars []byte

	l := len(input)

	for i := l - 1; i >= 0; i-- {
		chars = append(chars, input[i])
	}

	return string(chars)
}
