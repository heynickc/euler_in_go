package main

import (
	"fmt"
)

func main() {
	// divisors := []int{
	// 	10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	// }
	// for i := 1; ; i++ {
	// 	evenDivisors := make([]int, 0)
	// 	for _, num := range divisors {
	// 		factorOfN := i % num
	// 		if factorOfN == 0 {
	// 			evenDivisors = append(evenDivisors, num)
	// 		}
	// 	}
	// 	if len(evenDivisors) == 11 {
	// 		fmt.Printf("%v is divisible by %v\n", i, evenDivisors)
	// 		break
	// 	}
	// }

	// Optimization?
	// var factors = [...]int{2, 3, 2, 5, 2, 7, 3, 11, 13, 2, 17, 19} // GCD or LCM?
	// var num int = 1
	// for key, _ := range factors {
	// 	num *= factors[key]
	// }
	// fmt.Println(num)

	// fmt.Println(lcm(2, 3))
	fmt.Println(lcmm(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
}

// GCD
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	z := (x * y) / gcd(x, y)
	return z
}

// LCM can use GCD, so
// LCM (m, n) = abs((m * n))/gcd(m, n)
// https://stackoverflow.com/questions/147515/least-common-multiple-for-3-or-more-numbers/147539#147539

func lcmm(args ...int) []int {
	x := []int{lcm(args[0], args[1])}
	for i, num := range args {
		if i < len(args)-1 {
			x = append(x, lcm(x[i], num))
		}
	}
	return x
}
