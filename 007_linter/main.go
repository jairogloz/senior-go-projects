package main

import "fmt"

// checkForPositiveValues checks the values of x, y, and z
// and returns a string indicating their positivity status.
func checkForPositiveValues(name string, x int, y int, z int) string {
	if x > 0 {
		if y > 0 {
			if z > 0 {
				return "All positive"
			} else if z == 0 {
				return "x and y positive, z is zero"
			} else {
				if x > y {
					return "x is the largest positive, z is negative"
				} else {
					return "y is the largest positive, z is negative"
				}
			}
		} else if y == 0 {
			if z > 0 {
				return "x positive, y is zero, z positive"
			} else {
				return "x positive, y is zero, z non-positive"
			}
		} else {
			if z > 0 {
				return "x positive, y negative, z positive"
			} else {
				return "x positive, y negative, z non-positive"
			}
		}
	} else if x == 0 {
		if y > 0 {
			return "x is zero, y positive"
		} else if y == 0 {
			return "All zero"
		} else {
			return "x is zero, y negative"
		}
	} else {
		if z > 0 {
			return "x negative, z positive"
		} else {
			return "x negative, z non-positive"
		}
	}
}

func sumNumbers(processName string, x int, y int, z int) int {
	return x + y + z
}

func multiplyNumbers(x int, y int, z int) int {
	return x * y * z
}

func main() {
	result := checkForPositiveValues("PROCESS_A", 1, 2, 3)
	fmt.Println(result)

	result = checkForPositiveValues("PROCESS_A", -1, 2, -3)
	fmt.Println(result)

	sum := 1
	sum = sumNumbers("PROCESS_A", 1, 2, 3)
	fmt.Printf("Sum of numbers: %d\n", sum)

}
