package trial_division

import (
	"math/big"
	"testing"
)

type testPairBig struct {
	val *big.Int
	out bool
}

type testPairInt struct {
	val int
	out bool
}

var testsBig = []testPairBig{
	{big.NewInt(1), false},
	{big.NewInt(2), true},
	{big.NewInt(3), true},
	{big.NewInt(4), false},
	{big.NewInt(5), true},
	{big.NewInt(9), false},
	{big.NewInt(15), false},
	{big.NewInt(17), true},
	{big.NewInt(18), false},
	{big.NewInt(19), true},
	{big.NewInt(20), false},
	{big.NewInt(21), false},
	{big.NewInt(22), false},
	{big.NewInt(23), true},
	{big.NewInt(24), false},
	{big.NewInt(97), true},
	{big.NewInt(100), false},
	{big.NewInt(137), true},
	{big.NewInt(200), false},
	{big.NewInt(379), true},
}

var testsInt = []testPairInt{
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{5, true},
	{9, false},
	{15, false},
	{17, true},
	{18, false},
	{19, true},
	{20, false},
	{21, false},
	{22, false},
	{23, true},
	{24, false},
	{97, true},
	{100, false},
	{137, true},
	{200, false},
	{379, true},
}

func TestIsPrimeTrialDivisionBig(t *testing.T) {
	// fmt.Println("\nStart of Testing math.Big Trial Division Primality Testing...")
	for _, pair := range testsBig {
		result := IsPrimeTrialDivisionBig(pair.val)
		if result != pair.out {
			t.Errorf("Expected %v to be %v, but got %v", pair.val, pair.out, result)
		}
	}
}

func TestIsPrimeTrialDivision(t *testing.T) {
	// fmt.Println("\nStart of Testing Standard Library Trial Division Primality Testing...")
	for _, pair := range testsInt {
		result := IsPrimeTrialDivision(pair.val)
		if result != pair.out {
			t.Errorf("Expected %v to be %v, but got %v", pair.val, pair.out, result)
		}
	}
}
