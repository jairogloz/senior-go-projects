package main

import "fmt"

/*
Given an array of integers nums and an integer k, find the maximum sum of any k consecutive elements in
the array. Your solution should be as efficient as possible.
*/
func main() {
	nums := []int{2, 1, 5, 85, 1, 30, 2}
	k := 3
	fmt.Println(maxSum(nums, k))
}

func maxSum(num []int, k int) int {
	// []int{2, 1, 5, 1, 3, 2}
	//                k=3
	// mSum = 8
	// currentSum=8/6/7
	mSum := 0
	currentSum := 0
	for i := 0; i < k; i++ {
		mSum += num[i]
		currentSum = mSum
	}
	for i := k; i < len(num); i++ { // i = 3/4
		currentSum -= num[i-k]
		currentSum += num[i]
		if currentSum > mSum {
			mSum = currentSum
		}
	}
	return mSum
}
