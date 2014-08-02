// http://rosettacode.org/wiki/Fibonacci_sequence

package iterative

import (
	"math/big"
)

func IterativeFib(n uint64) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}
