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
	var s = "In girum imus nocte et consumimur igni"
	var sNum = "90099009"
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(sNum))

	var startNum = 999
	for i := startNum; i > 0; i-- {
		for i := 999; i > 0; i-- {
			if isPalindrome(strconv.FormatInt(int64(i*startNum), 10)) {
				fmt.Println(i * startNum)
				return
			}
		}
		startNum -= 1
	}
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
