package main

import (
	"math"
	"strconv"
	"strings"
)

func LastDigit(n1, n2 string) int {
	num1 := 0
	num2 := 0

	pattern := map[int][]int{
		0: {0},
		1: {1},
		2: {2, 4, 8, 6},
		3: {3, 9, 7, 1},
		4: {4, 6},
		5: {5},
		6: {6},
		7: {7, 9, 3, 1},
		8: {8, 4, 2, 6},
		9: {9, 1},
	}

	if strings.Contains(n1, "^") {
		pow1 := strings.Split(n1, "^")
		a1, _ := strconv.Atoi(pow1[0])
		a2, _ := strconv.Atoi(pow1[1])

		pow := math.Pow(float64(a1), float64(a2))
		stringPow := strconv.Itoa(int(pow))
		last := string(stringPow[len(stringPow)-1])
		num1, _ = strconv.Atoi(last)

	} else {
		num1, _ = strconv.Atoi(string(n1[len(n1)-1]))
	}

	if strings.Contains(n2, "^") {
		pow2 := strings.Split(n2, "^")
		a1, _ := strconv.Atoi(pow2[0])
		a2, _ := strconv.Atoi(pow2[1])
		num2 = int(math.Pow(float64(a1), float64(a2)))

	} else {
		num2, _ = strconv.Atoi(n2)
	}

	if len(pattern[num1]) == 0 {
		return 0
	}

	res := (num2 - 1) % len(pattern[num1])
	return pattern[num1][res]

}

