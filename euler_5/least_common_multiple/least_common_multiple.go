package least_common_multiple

// GCD
func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// LCM can use GCD, so
// LCM (m, n) = abs((m * n))/gcd(m, n)
// https://stackoverflow.com/questions/147515/least-common-multiple-for-3-or-more-numbers/147539#147539

// LCM
func LCM(x, y int) int {
	z := (x * y) / GCD(x, y)
	return z
}

// LCM MULTIPLE INTEGERS
func LCMM(factors []int) []int {
	x := []int{2}
	for i, num := range factors {
		if i < len(factors) {
			x = append(x, LCM(x[i], num))
		}
	}
	// x, y := factors[0], 0
	// for _, num := range factors {
	// 	x, y = y, LCM(x, num)
	// }
	// return x
	return x
}
