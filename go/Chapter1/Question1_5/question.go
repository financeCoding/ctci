package main

import "fmt"
import "strconv"
import "strings"

func countCompression(str string) int {
	last := rune(str[0])
	size := 0
	count := 1
	for _, c := range str {
		if c == last {
			count++
		} else {
			last = c
			size += 1 + len(strconv.Itoa(count))
			count = 1
		}
	}

	size += 1 + len(strconv.Itoa(count))
	return size
}

func compressAlternate(str string) string {
	size := countCompression(str)
	if size >= len(str) {
		return str
	}

	array := make([]string, size)
	index := 0
	last := rune(str[0])
	count := 1
	for i, c := range str {
		if i == 0 {
			continue // skip the first element in the string. 
		}
		if c == last {
			count++
		} else {
			index = setChar(array, last, index, count)
			last = c
			count = 1
		}
	}

	index = setChar(array, last, index, count)
	return strings.Join(array, "")
}

func setChar(array []string, last rune, index int, count int) int {
	array[index] = string(last)
	index++
	cnt := strings.Split(strconv.Itoa(count), "")
	for _, c := range cnt {
		array[index] = string(c)
		index++
	}

	return index
}

func main() {
	str := "abbccccccde"
	c := countCompression(str)
	str2 := compressAlternate(str)
	fmt.Println("Old String (len = %d): %s", len(str), str)
	fmt.Println("New String (len = %d): %s", len(str2), str2)
	fmt.Println("Potential Compression: %d", c)
}
