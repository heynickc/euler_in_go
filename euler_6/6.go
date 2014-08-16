// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385

// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025

// Hence the difference between the sum of the squares of the
// first ten natural numbers and the square of the sum is
// 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the
// first one hundred natural numbers and the square of the sum.

package main

import (
	"fmt"
	"math"
)

func main() {
	var totesSquare float64
	var squareTotes float64
	for i := 1; i < 100; i++ {
		totesSquare += math.Pow(float64(i), float64(2))
		squareTotes += float64(i)
	}
	squareTotes = math.Pow(squareTotes, 2)

	fmt.Println(totesSquare)
	fmt.Println(squareTotes)
	fmt.Printf("Difference %v - %v = %f\n", squareTotes, totesSquare, (squareTotes - totesSquare))
}
