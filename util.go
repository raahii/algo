package algo

func BinComb(n int, c chan []int) {
	n2 := 1 << uint(n)

	for i := 0; i < n2; i++ {
		pat := make([]int, n)
		for j := 0; j < n; j++ {
			pat[j] = i >> uint(j) & 1
		}
		c <- pat
	}
	close(c)
}
