package main

import (
	"strconv"
	"unicode"
)

func IncrementString(strng string) string {
	if len(strng) == 0 {
		return strng
	}

	index := len(strng) - 1
	for index >= 0 && unicode.IsDigit(rune(strng[index])) {
		index--
	}

	head := strng[:index+1]
	tail := strng[index+1:]

	if len(tail) == 0 {
		return head + "1"
	}

	num, _ := strconv.Atoi(tail)
	num++
	res := strconv.Itoa(num)
	return head + res

}

