package main

import (
	"bufio"
	"fmt"
	"os"
)

/* IO */

func ReadInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func ReadInts(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	return nums
}

var sc = bufio.NewScanner(os.Stdin)

func ReadLine() string {
	sc.Scan()
	return sc.Text()
}

func ReadWords(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&strs[i])
	}
	return strs
}
