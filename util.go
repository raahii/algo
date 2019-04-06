package algo

func AllComb(n int) [][]int {
	n2 := 1 << uint(n)
	ret := make([][]int, n2)

	for i := 0; i < n2; i++ {
		pat := make([]int, n)
		for j := 0; j < n; j++ {
			pat[j] = i >> uint(j) & 1
		}
		ret[i] = pat
	}
	return ret
}

func AllCombChan(n int, c chan []int) {
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

func PermuteInts(nums []int) [][]int {
	n := len(nums)
	nPatterns := Factorial(n)

	ret := make([][]int, 0, nPatterns)
	var cur []int
	permuteInts(n, cur, nums, &ret)

	return ret
}

func permuteInts(n int, cur, nums []int, ret *[][]int) {
	if n == 0 {
		*ret = append(*ret, cur)
		return
	}

	for i := 0; i < n; i++ {
		rest := make([]int, 0, n-1)
		for j := 0; j < n; j++ {
			if i != j {
				rest = append(rest, nums[j])
			}
		}
		permuteInts(n-1, append(cur, nums[i]), rest, ret)
	}
}
