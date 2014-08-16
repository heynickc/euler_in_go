package prime_factorization

func LCMPrimeFactors() int {
	// Optimization with Prime Factorization?
	// https://github.com/hermanschaaf/go-euler/blob/master/5.go

	factors := []int{2, 3, 2, 5, 2, 7, 3, 11, 13, 2, 17, 19}
	var num int = 1
	for _, factor := range factors {
		num *= factor
	}
	return num
}
