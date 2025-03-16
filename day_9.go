package main

import "fmt"

func Decode(roman string) int {
	data := make(map[rune]int)
	res := 0
	length := len(roman)

	data['I'] = 1
	data['V'] = 5
	data['X'] = 10
	data['L'] = 50
	data['C'] = 100
	data['D'] = 500
	data['M'] = 1000

	for i, char := range roman {
		value := data[char]
        if i < length-1 && value < data[rune(roman[i+1])] {
            res -= value 
        } else {
            res += value 
        }
	}

	return res
}

func main() {
	fmt.Println(Decode("IX"))
}
