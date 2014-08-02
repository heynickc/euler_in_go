// http://rosettacode.org/wiki/Primality_by_trial_division#Go

package trial_division

import (
	"math"
	"math/big"
)

func IsPrimeTrialDivisionBig(p *big.Int) bool {
	// less than 2 not prime
	if p.Cmp(big.NewInt(2)) == -1 {
		return false
	}

	// 2 is prime
	if p.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	// Even numbers are not prime
	// Compares Mod of p%2 to 0 using Cmp method
	// if this returns 0, then the numbers are equal
	if big.NewInt(0).Cmp(new(big.Int).Mod(p, big.NewInt(2))) == 0 {
		return false
	}

	// Starting at 3, check odd numbers
	// only check up to the square root
	// so stop when i <= square root of p

	for i := big.NewInt(3); new(big.Int).Mul(i, i).Cmp(p) != 1; i.Add(i, big.NewInt(2)) {
		modP := new(big.Int).Mod(p, i)
		// log.Printf("%v / %v big.Int with remainder %v", p, i, modP)
		if modP.Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}
	return true
}

func IsPrimeTrialDivision(p int) bool {
	if p < 2 {
		return false
	}
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	pAsFloat64 := float64(p)
	for i := 3; float64(i) <= math.Sqrt(pAsFloat64); i += 2 {
		// log.Printf("%v / %v = %v remainder %v", p, i, p/i, p%i)
		if p%i == 0 {
			return false
		}
	}
	return true
}
