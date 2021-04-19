package utils

import "unicode"

// This function returns true when the input strings are anagrams of one another.
func IsAnagram(input1, input2 string) bool {

	map1 := getMap(input1)
	map2 := getMap(input2)

	if len(map1) != len(map2) {
		return false
	}

	isAnagram := true

	for key, c1 := range map1 {

		if !isAnagram {
			return false //quick exit
		}

		map2Entry := map2[key]

		if map2Entry == c1 {
			//this character has the same number of occurrences...
			//could be an anagram
			isAnagram = true
			continue
		}

		if map2Entry == 0 {
			isAnagram = false
			continue
		}

		if map2Entry != c1 {
			isAnagram = false
			continue
		}
	}

	return isAnagram
}

// Convert the input string into a map[rune]int
// Each entry in the map provides a count of the instances of that character in the string
func getMap(input string) map[rune]int {
	m := map[rune]int{}

	for _, c := range input {

		if unicode.IsSpace(c) {
			continue
		}

		key := unicode.ToLower(c)

		if m[key] == 0 {
			m[key] = 1
		} else {
			m[key] += 1
		}
	}

	return m
}
