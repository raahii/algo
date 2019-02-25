package algo

import (
	"bufio"
	"fmt"
	"os"
)

/* IO */

func ReadInt(f *os.File) int {
	var n int
	fmt.Fscanf(f, "%d", &n)
	return n
}

func ReadInts(f *os.File, n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(f, "%d", &nums[i])
	}
	return nums
}

func ScanInts(f *os.File, vars ...*int) {
	for i := 0; i < len(vars); i++ {
		fmt.Fscanf(f, "%d", vars[i])
	}
}

func ReadLines(f *os.File, n int) []string {
	var lines []string
	reader := bufio.NewReader(f)
	for i := 0; i < n; i++ {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(bytes))
	}

	return lines
}

func ReadWords(f *os.File, n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(f, "%s", &strs[i])
	}
	return strs
}
