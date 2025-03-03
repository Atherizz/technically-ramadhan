package main

import "fmt"

func CountPositivesSumNegatives(numbers []int) []int {
	positiveCount, negativeSum := 0,0
	if len(numbers) == 0 {
		return []int{}
	} else {
		for _, number := range(numbers) {
			if number > 0 {
				positiveCount++
			} else if number < 0 {
				negativeSum += number
			}
		}
		return []int{positiveCount, negativeSum}
	}
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	result := CountPositivesSumNegatives(numbers)
	fmt.Println(result)
}
