package main

import "fmt"

// optimal
func isUniqueChars(str string) bool {
	if len(str) > 256 {
		return false
	}

	checker := 0

	for i := 0; i < len(str); i++ {
		val := str[i] - 'a'
		if (checker & (1 << val)) > 0 {
			return false
		} else {
			checker |= (1 << val)
		}
	}

	return true
}

// optimal
func isUniqueChars2(str string) bool {
	if len(str) > 256 {
		return false
	}

	char_set := make([]bool, 256)

	for i := 0; i < len(str); i++ {
		val := str[i]
		if char_set[val] != false {
			return false
		} else {
			char_set[val] = true
		}
	}

	return true
}

// least optimal (entire string is iterated on)
//func isUniqueChars3(str string) bool {
// Sets dont really exist in go, just do the map example.
//}

// least optimal
func isUniqueChars4(str string) bool {
	if len(str) > 256 {
		return false
	}

	m := make(map[uint8]uint8)
	for i := 0; i < len(str); i++ {
		_, isContained := m[str[i]]
		if isContained {
			return false
		}

		m[str[i]] = str[i]
	}

	return true
}

func main() {
	fmt.Println("hello world")

	words := []string{"abcde", "hello", "apple", "kite", "padle"}

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s %t %t %t\n",
			words[i],
			isUniqueChars(words[i]),
			isUniqueChars2(words[i]),
			isUniqueChars4(words[i]))
	}
}
