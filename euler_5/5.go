package main

import (
	"fmt"
)

func main() {
	lowNums := []int{
		2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	for i := 1; i < 3000; i++ {
		evenDivisors := make([]int, 0)
		for _, num := range lowNums {
			factorOfN := i % num
			if factorOfN == 0 {
				evenDivisors = append(evenDivisors, num)
			}
			if len(evenDivisors) > 5 {
				fmt.Printf("%v is divisible by %v\n", i, evenDivisors)
			}
		}
	}
}
