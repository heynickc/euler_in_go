package recursive

func RecursiveFib(a int) int {
	if a < 2 {
		return a
	}
	return RecursiveFib(a-1) + RecursiveFib(a-2)
}
