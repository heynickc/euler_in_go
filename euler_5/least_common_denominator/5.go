package main

import (
	"fmt"
)

func main() {
	fmt.Println(lcmm(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
}

// GCD
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// LCM
func lcm(x, y int) int {
	z := (x * y) / gcd(x, y)
	return z
}

// LCM can use GCD, so
// LCM (m, n) = abs((m * n))/gcd(m, n)
// https://stackoverflow.com/questions/147515/least-common-multiple-for-3-or-more-numbers/147539#147539

// LCM MULTIPLE
func lcmm(args ...int) []int {
	x := []int{lcm(args[0], args[1])}
	for i, num := range args {
		if i < len(args)-1 {
			x = append(x, lcm(x[i], num))
		}
	}
	return x
}
