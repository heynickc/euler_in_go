package main

import (
	"./iterative"
	// "./recursive"
	"fmt"
	"math/big"
)

func main() {
	// Recursive
	// a := 50
	// var x []int
	// for i := 0; i < a; i++ {
	// 	if recursive.RecursiveFib(i)%2 == 0 {
	// 		x = append(x, recursive.RecursiveFib(i))
	// 	}
	// }
	// fmt.Println(x)

	// Interative
	max, result, totes := uint64(4000000), big.NewInt(0), big.NewInt(0)
	for i := uint64(0); result.Uint64() < max; i++ {
		result = iterative.IterativeFib(i)
		modFib := new(big.Int).Mod(result, big.NewInt(2))
		if modFib.Cmp(big.NewInt(0)) == 0 {
			totes.Add(totes, result)
		}
	}
	fmt.Println(totes)

}
