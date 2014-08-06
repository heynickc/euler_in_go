package main

import (
	"fmt"
)

func main() {
	divisors := []int{
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}
	for i := 1; ; i++ {
		evenDivisors := make([]int, 0)
		for _, num := range divisors {
			factorOfN := i % num
			if factorOfN == 0 {
				evenDivisors = append(evenDivisors, num)
			}
		}
		if len(evenDivisors) == 11 {
			fmt.Printf("%v is divisible by %v\n", i, evenDivisors)
			break
		}
	}
}
