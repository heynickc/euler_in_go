// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "In girum imus nocte et consumimur igni"
	var sNum = "90099009"
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(sNum))
}

func isPalindrome(str string) bool {
	s := strings.Replace(strings.ToLower(str), " ", "", -1)
	mid := len(s) / 2  // loop stops at the middle
	last := len(s) - 1 // starting character to match against
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			// compare the index at i to the
			// index at the last character minus i
			return false
		}
	}
	return true
}
