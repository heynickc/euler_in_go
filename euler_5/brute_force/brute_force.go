package brute_force

import (
	"fmt"
)

func LCM(factors []int) int {
	var result int
	for i := 1; ; i++ {
		evenDivisors := make([]int, 0)
		for _, num := range factors {
			factorOfN := i % num
			if factorOfN == 0 {
				evenDivisors = append(evenDivisors, num)
			}
		}
		if len(evenDivisors) == 20 {
			fmt.Printf("%v is divisible by %v\n", i, evenDivisors)
			result = i
			break
		}
	}
	return result
}
