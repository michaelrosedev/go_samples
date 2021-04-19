package utils

// Removes any map entries with a value greater than 1
func ToUniqueKeys(input map[string]int) map[string]int {
	for k, v := range input {
		if v > 1 {
			delete(input, k)
		}
	}

	return input
}
