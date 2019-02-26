package algo

func Pow(a, n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c *= a
	}
	return c
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
