// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

// http://rosettacode.org/wiki/Prime_decomposition

// If a number n is not a prime, it can be factored into two factors a and b:
// n = a*b
// If both a and b were greater than the square root of n,
// a*b would be greater than n. So at least one of those factors
// must be less or equal to the square root of n,
// and to check if n is prime, we only need to test for
// factors less than or equal to the square root.

// Define the number
// List all of the prime numbers below that number
// Find the lowest prime number that goes into that number
// If the prime number goes into the number, then that is the first prime factor
// Add that first prime factor to an array
// Grab the result of that division
// Test the result of that division for the lowest prime number that goes into it
// Add that number to the prime factor array
// Grab the result of that division
// Continue until the result of the division is a prime number
// Add that number to the array
// Print the array

package main

import (
	// "./miller_rabin"
	"./trial_division"
	"fmt"
	// "math/big"
	// "strconv"
)

func main() {
	var num int = 600851475143
	for i := 2; i < num/2; i++ {
		// only checks up to half of the number, reason below
		if num%i == 0 { // if the number is a factor of num, enter
			fmt.Printf("%v is a factor of %v \n", i, num)
			var pair int = num / i
			// the first prime factor's multiple is the largest prime factor
			// the first prime result is what the LARGEST PRIME FACTOR
			// is multiplied by to get the SMALLEST PRIME FACTOR
			if trial_division.IsPrimeTrialDivision(pair) {
				fmt.Printf("%v is the largest prime factor \n", num)
				fmt.Printf("%v is the small pair to largest prime factor \n", pair)
				fmt.Println(pair)
				return
			}
		}
	}
}
