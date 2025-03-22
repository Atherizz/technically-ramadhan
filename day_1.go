package main



func CountPositivesSumNegatives(numbers []int) []int {
	positiveCount, negativeSum := 0, 0
	if len(numbers) == 0 {
		return []int{}
	} else {
		for _, number := range numbers {
			if number > 0 {
				positiveCount++
			} else if number < 0 {
				negativeSum += number
			}
		}
		return []int{positiveCount, negativeSum}
	}
}


