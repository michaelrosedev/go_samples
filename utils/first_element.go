package utils

import "sort"

// This function returns the first unique string from a slice of strings.
//
// The "first unique element" means the first element for which there is only
// 1 matching value in the slice. The term "first" means the first string ordered alphabetically.
func FirstUniqueElement(input []string) string {
	m := make(map[string]int)

	for _, s := range input {
		if m[s] == 0 {
			m[s] = 1
		} else {
			m[s] += 1
		}
	}

	m = ToUniqueKeys(m)

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	if len(keys) > 0 {
		return keys[0]
	} else {
		return ""
	}
}
