package algo

func Pow(a, n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c *= a
	}
	return c
}
