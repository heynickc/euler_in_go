package miller_rabin

import (
	"math/big"
)

func IsPrimeMillerRabinCollection(primes []string) []*big.Int {
	validNums := make([]*big.Int, 0)
	for _, s := range primes {
		p, _ := new(big.Int).SetString(s, 10)
		if IsPrimeMillerRabin(p) {
			validNums = append(validNums, p)
		}
	}
	return validNums
}

func IsPrimeMillerRabin(p *big.Int) bool {
	nreps := 20
	if p.ProbablyPrime(nreps) {
		return true
	} else {
		return false
	}
}
