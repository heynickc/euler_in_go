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
	var factors = [...]int{2, 3, 2, 5, 2, 7, 3, 11, 13, 2, 17, 19} // GCD or LCM?
	var num int = 1
	for key, _ := range factors {
		num *= factors[key]
	}
	fmt.Println(num)
}

// GCD
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// LCM can use GCD, so
// LCM (m, n) = abs((m * n))/gcd(m, n)
