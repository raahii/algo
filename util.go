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
	n := Factorial(len(nums))
	ret := make([][]int, 0, n)
	permuteInts(nums, &ret)
	return ret
}

func permuteInts(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}
