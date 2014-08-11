package least_common_multiple

// LCM can use GCD, so
// LCM (m, n) = abs((m * n))/gcd(m, n)
// https://stackoverflow.com/questions/147515/least-common-multiple-for-3-or-more-numbers/147539#147539

// GCD
func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// LCM
func LCM(x, y int) int {
	z := (x * y) / GCD(x, y)
	return z
}

// LCM MULTIPLE INTEGERS
func LCMM(factors []int) int {
	var x, y int = factors[0], factors[1]
	for i := 1; i < len(factors); i++ {
		x, y = LCM(x, y), factors[i]
	}
	return x
}
