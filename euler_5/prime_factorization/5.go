package main

import (
	"fmt"
)

func main() {
	// Optimization with Prime Factorization?
	// https://github.com/hermanschaaf/go-euler/blob/master/5.go

	var factors = [...]int{2, 3, 2, 5, 2, 7, 3, 11, 13, 2, 17, 19} // GCD or LCM?
	var num int = 1
	for key, _ := range factors {
		num *= factors[key]
	}
	fmt.Println(num)
}
