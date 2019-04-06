package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func findFunction(str string) (string, string, error) {
	startPos := strings.Index(str, "func")
	if startPos == -1 {
		return "", "", fmt.Errorf("Variable 'str' doesn't have a function declaration")
	}

	offset, nOpens := startPos, 0
	iter, maxIter := 0, 100000
	for {
		posOpen := strings.Index(str[offset:], "{")
		posClose := strings.Index(str[offset:], "}")

		if posClose == -1 {
			return "", "", fmt.Errorf("Closing brace of function is not found!")
		}

		if posOpen == -1 || posClose < posOpen {
			nOpens--
			if nOpens == 0 {
				return str[startPos : offset+posClose+2], str[offset+posClose+2:], nil
			}
			offset += posClose + 1
		} else {
			nOpens++
			offset += posOpen + 1
		}

		iter++
		if iter > maxIter {
			return "", "", fmt.Errorf("Exceeded maximum iterations, %d", maxIter)
		}
	}
}

func main() {
	// 下記で指定したファイルの中から関数部分だけを
	// 抜き出してきてvimにコピペできる文字列を出力する
	files := []string{
		"../array.go",
		"../io.go",
		"../calc.go",
		"../util.go",
	}

	allFuncStr := ""
	for _, filename := range files {
		// read file content
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		content := string(bytes)

		// find function strings
		for {
			funcStr, restContent, err := findFunction(content)
			if err != nil {
				break
			}
			allFuncStr += funcStr + "\n"
			content = restContent
		}
	}

	// add comment and credit
	credit := "// source: https://github.com/raahii/algo\n"
	allFuncStr = "// {{{\n" + credit + allFuncStr + "// }}}\n"
	fmt.Println(allFuncStr)
}
