package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "perro viejo paracaidista"

	fmt.Println("reversString: ", reverseString(s))
	fmt.Println("removeDuplicates: ", removeDuplicates(s))
	fmt.Println("firstUniqueChar: ", firstUniqueChar(s))

}

func reverseString(s string) string {
	sb := strings.Builder{}

	for j := len(s) - 1; j >= 0; j-- {
		sb.WriteRune(rune(s[j]))
	}
	return sb.String()
}

func removeDuplicates(s string) string {
	seen := map[rune]bool{}
	newString := []rune{}
	for _, r := range s {
		if seen[r] {
			continue
		}
		seen[r] = true
		newString = append(newString, r)
	}
	return string(newString)
}

func firstUniqueChar(s string) string {
	// "perro viejo paracaidista"
	var unique rune
	count := map[rune]int{}
	for _, r := range s {
		count[r] = count[r] + 1
	}
	for _, char := range s {
		if count[char] == 1 {
			return string(char)
		}
	}
	return string(unique)
}
