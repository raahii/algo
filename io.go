package algo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/* IO */

func ReadInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func ReadIntF(f io.Reader) int {
	var n int
	fmt.Fscanf(f, "%d", &n)
	return n
}

func ReadInts(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	return nums
}

func ReadIntsF(f io.Reader, n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(f, "%d", &nums[i])
	}
	return nums
}

func ReadWord(f io.Reader) string {
	var str string
	fmt.Scan(&str)
	return str
}

func ReadWordF(f io.Reader) string {
	var str string
	fmt.Fscanf(f, "%s", &str)
	return str
}

func ReadLine() string {
	r := bufio.NewReader(os.Stdin)
	str, _ := r.ReadString('\n')
	str = strings.Trim(str, "\n")
	return str
}

func ReadLineF(f io.Reader) string {
	r := bufio.NewReader(f)
	str, _ := r.ReadString('\n')
	str = strings.Trim(str, "\n")
	return str
}

func ReadLines(n int) []string {
	r := bufio.NewReader(os.Stdin)
	lines := make([]string, n)
	for i := 0; i < n; i++ {
		lines[i], _ = r.ReadString('\n')
		lines[i] = strings.Trim(lines[i], "\n")
	}
	return lines
}

func ReadLinesF(f io.Reader, n int) []string {
	r := bufio.NewReader(f)
	lines := make([]string, n)
	for i := 0; i < n; i++ {
		lines[i], _ = r.ReadString('\n')
		lines[i] = strings.Trim(lines[i], "\n")
	}
	return lines
}
