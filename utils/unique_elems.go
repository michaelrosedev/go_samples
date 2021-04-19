package utils

import "sort"

// Identifies all string(s) that occur only once in the input and returns a new slice containing the unique elements only
func UniqueEntries(input []string) []string {

	var m = make(map[string]int)
	for _, s := range input {
		if m[s] == 0 {
			m[s] = 1
		} else {
			m[s] += 1
		}
	}

	m = ToUniqueKeys(m)

	var res []string

	for k := range m {
		res = append(res, k)
	}

	sort.Strings(res)

	return res
}
