package main

import (
	// "./brute_force"
	"./least_common_multiple"
	"fmt"
)

func main() {
	// factors := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	factors := []int{1, 2, 3}

	// fmt.Println(brute_force.LCM(factors))
	fmt.Println(least_common_multiple.LCMM(factors))
}
