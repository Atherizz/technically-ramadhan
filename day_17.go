package main

import (
	"math"
)

func Buddy(start, limit int) []int {
	for i := start; i <= limit; i++ {
		sumDivI := Divisors(i) - 1
		if sumDivI > i {
			sumDivJ := Divisors(sumDivI) - 1
			if sumDivJ == i {
				return []int{sumDivJ, sumDivI}
			}
		}
	}

	return []int{}
}

func Divisors(n int) int {
	sum := 0
	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			sum += i
			if i != 1 && i != n/i {
				sum += n / i
			}
		}
	}
	return sum
}

