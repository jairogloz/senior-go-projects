package main

import (
	"fmt"
	"log"
)

// Problem Description:
//
// You are tasked with implementing a function
// CalculateFactorial in Go that calculates the
// factorial of a non-negative integer. The
// factorial of a non-negative integer n is the
// product of all positive integers less than or
// equal to n.
//
// Output: 120 (5! = 5 x 4 x 3 x 2 x 1 = 120)
// result := 20
// result = result x i 
// Output: 1 (0! = 1)
// Output: 6 (3! = 3 x 2 x 1 = 6)
//
// Constraints:
//
// The input integer n is non-negative (0 <= n <= 10).
func main() {
	fmt.Println(CalculateFactorialRecursive(5))
	fmt.Println(CalculateFactorialRecursive(0))
	fmt.Println(CalculateFactorialRecursive(3))
}

func CalculateFactorialRecursive(n int) int {
	if n < 0 {
		log.Println("factorial not defined for negative numbers")
		return 0
	}
	if n > 10 {
		log.Println("n must be less than or equal to 10")
		return 0
	}
	if n == 0 {
		return 1 // 0! is defined as 1
	}

	result := n  * CalculateFactorialRecursive(n-1)
	
	return result
}

func CalculateFactorial(n int) int {
	if n < 0 {
		log.Println("factorial not defined for negative numbers")
		return 0
	}
	if n > 10 {
		log.Println("n must be less than or equal to 10")
		return 0
	}

	result := 1
	for i:=n;i>=1;i-- {
		result = result * i
	}
	
	return result
}
