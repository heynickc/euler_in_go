// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var max = 0
	for i := 999; i > 0; i-- {
		for n := 999; n > 0; n-- {
			if isPalindrome(strconv.FormatInt(int64(i*n), 10)) {
				if i*n > max {
					max = i * n
				}
			}
		}
		// https://github.com/hermanschaaf/go-euler/blob/master/4.go
		// fmt.Println(1100*i - 990*(i/10) - 99*(i/100))
	}
	fmt.Println(max)
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
