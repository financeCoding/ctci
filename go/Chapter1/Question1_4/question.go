package main

import "fmt"

func replaceSpaces(str []rune, length, spacesCount, addtionalSizePerSpace int) {
	index := length + spacesCount*addtionalSizePerSpace

	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[index-1] = '0'
			str[index-2] = '2'
			str[index-3] = '%'
			index = index - 3
		} else {
			str[index-1] = str[i]
			index--
		}
	}
}

func main() {
	str := "abc d e f"
	spacesCount := 0
	addtionalSizePerSpace := 2

	// Count the number of spaces to allocate space for
	for _, c := range str {
		if c == ' ' {
			spacesCount++
		}
	}

	// Removed +1 from [arr] since we dont need null terminating strings here
	arr := make([]rune, len(str)+spacesCount*addtionalSizePerSpace)

	for i, c := range str {
		arr[i] = c
	}

	replaceSpaces(arr, len(str), spacesCount, addtionalSizePerSpace)
	fmt.Println(string(arr))
}
