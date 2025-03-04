package main

func WordsToMarks(s string) int {
	res := 0
	for _, letter := range s {
		val := int(letter) - int('a') + 1
		res += val
	}
	return res
}




