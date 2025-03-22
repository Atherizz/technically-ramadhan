package main

var fibo = make([]int, 100)

func Perimeter(n int) int {
	return 4 * fibonacci(n+1)
}

func fibonacci(n int) int {
	fibo[0] = 1
	fibo[1] = 1
	sum := 2

	for i := 2; i < n; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
		sum += fibo[i]
	}

	return sum

}
