package main

import "fmt"
import "sort"
import "strings"

func permutation(s, t string) bool {
	tmp_s, tmp_t := strings.Split(s, ""), strings.Split(t, "")
	sort.Strings(tmp_s)
	sort.Strings(tmp_t)
	return strings.Join(tmp_s, "") == strings.Join(tmp_t, "")
}

func main() {
	words := [][]string{{"apple", "papel"},
		{"carrot", "tarroc"},
		{"hello", "llloh"}}

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s %t\n", words[i], permutation(words[i][0], words[i][1]))
	}
}
