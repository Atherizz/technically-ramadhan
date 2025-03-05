package main

import "strings"


func toWeirdCase(str string) string {
	separated := strings.Split(str, " ")
	var res string
	for i, word := range separated {
		for j, letter := range word {
			char := string(letter)
			if j%2 == 0 {
				res += strings.ToUpper(char)
			} else {
				res += strings.ToLower(char)
			}
		}
		if i != (len(separated) - 1) {
			res += " "
		}
	}
	return res
}
