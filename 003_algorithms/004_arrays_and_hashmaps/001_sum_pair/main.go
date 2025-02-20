package main

import "fmt"

/*
Given a sorted array of integers nums and a target integer target, find two numbers in the array that add up to target. Return the indices of these two numbers (1-based index). You must use the two-pointer technique to solve this efficiently.
Constraints:

	The input array is sorted in ascending order.
	There is exactly one unique solution.
	The solution should run in O(n) time complexity.
	You cannot use the same element twice.
*/
func main() {
	nums := []int{1, 2, 3, 6, 8, 11, 15}
	target := 15
	sumPair(nums, target)
}

func sumPair(nums []int, target int) {
	if len(nums) == 0 {
		fmt.Println("empty array")
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			fmt.Println(nums[left], "+", nums[right], "=", target)
			return
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	fmt.Println("no pairs")
}
