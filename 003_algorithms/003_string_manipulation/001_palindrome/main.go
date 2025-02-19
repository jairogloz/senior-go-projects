package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("madam"))
	fmt.Println(isPalindrome("anitalavalatina"))
	fmt.Println(isPalindrome("dogiscrucial"))
}

func isPalindrome(s string) bool {
	// racecar
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
