package main

import "fmt"

// CalculateInvestment calculates the future value of an investment and total interest earned.
// This function takes four parameters:
// 1. initialInvestment (float64): The initial amount of money invested.
// 2. annualDeposit (float64): The amount deposited into the investment each year.
// 3. years (int): The number of years the investment is held.
// 4. annualInterestRate (float64): The annual interest rate as a percentage (e.g., 5 for 5%).
// It returns two float64 values:
// - The first return value is the total value of the investment at the end of the specified period.
// - The second return value is the total interest earned over the investment period.
// The function assumes that the annual deposit is added at the start of each year and
// calculates interest based on the average value of the investment over the year.
func CalculateInvestment(initialInvestment, annualDeposit float64, years int, annualInterestRate float64) (float64, float64) {
	totalValue := initialInvestment
	totalInterestEarned := 0.0

	for year := 1; year <= years; year++ {
		interestEarned := (totalValue + (annualDeposit / 2)) * annualInterestRate / 100
		totalValue += annualDeposit + interestEarned
		totalInterestEarned += interestEarned
	}

	return totalValue, totalInterestEarned
}

func main() {
	finalValue, totalInterest := CalculateInvestment(10000, 5000, 10, 5)
	fmt.Printf("Final Investment Value: $%.2f\n", finalValue)
	fmt.Printf("Total Interest Earned: $%.2f\n", totalInterest)
}
