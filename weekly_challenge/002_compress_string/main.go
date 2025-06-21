package main

import (
	"fmt"
	"strings"
)

/*
Please create a function in the language of your choosing that compresses a string input in the following manner. Be sure to talk through your thought process, especially around defining and testing edge cases as you go.
Example input:


aaaabbaccc
Example output:


a4b2a1c3

abcd

a1b1c1d1
*/

func main() {

	s := "aaaabbaccc"

	fmt.Println(compress(s))

}

func compress(s string) string {
	if len(s) == 0 {
		return ""
	}

	sb := strings.Builder{}
	current := ""
	count := 0
	runeArray := []rune(s)
	current = string(runeArray[0])
	for _, c := range runeArray {
		if string(c) != current {
			// compress
			sb.WriteString(fmt.Sprintf("%s%d", current, count))
			// reset count current
			count = 0
			current = string(c)
		}
		count++
	}
	// end case
	sb.WriteString(fmt.Sprintf("%s%d", current, count))
	return sb.String()
}
