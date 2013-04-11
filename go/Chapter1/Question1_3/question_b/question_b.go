package main

import "fmt"

func permutation(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make([]int, 256)

	for _, c := range s {
		// Count number of each char in s
		letters[c]++
	}

	for _, c := range t {
		if letters[c] == 0 {
			letters[c] = -1
		} else {
			letters[c]--
		}

		if letters[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	words := [][]string{{"apple", "papel"},
		{"carrot", "tarroc"},
		{"hello", "llloh"}}

	for _, word := range words {
		fmt.Printf("%s %t\n", word, permutation(word[0], word[1]))
	}
}
