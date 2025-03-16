package main

import "fmt"

func luckCheck(arr []int) bool {
	left, right := []int{}, []int{}
	mid := (len(arr) / 2)
	if len(arr)%2 == 0 {
		left = arr[:mid]
		right = arr[mid:]
	} else {
		left = arr[:mid]
		right = arr[mid+1:]
	}
	
	if sumArr(left, 0, len(left)-1) == sumArr(right, 0, len(right)-1) {
		return true
	}

	return false
}

func sumArr(arr []int, l int, r int) int {
	if l == r {
		return arr[l]
	} else {
		mid := (l+r) / 2
		leftSum := sumArr(arr, l, mid)
		rightSum := sumArr(arr, mid+1, r)
		return leftSum + rightSum
	}
}
