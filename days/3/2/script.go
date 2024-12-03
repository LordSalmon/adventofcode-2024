package day03_02

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	content := utils.ReadAOCInput(3, 1)
	fmt.Println("sum of statements: ", sumMulStatements(getMulStatements((content))))
}

func getMulStatements(content string) [][]int {
	var out [][]int
	isEnabled := true
	for strings.Contains(content, "mul(") {
		startIndex := strings.Index(content, "mul(")
		contentBefore := content[:startIndex]
		latestDoOccurrence := strings.LastIndex(contentBefore, "do()")
		latestDontOccurrence := strings.LastIndex(contentBefore, "don't()")
		if latestDoOccurrence != -1 || latestDontOccurrence != -1 {
			isEnabled = latestDoOccurrence > latestDontOccurrence
		}
		if isEnabled {
			mulContent := content[startIndex+4 : strings.Index(content[startIndex+4:], ")")+startIndex+4]
			if strings.Contains(mulContent, ",") {
				values := strings.Split(mulContent, ",")
				if len(values) != 2 || strings.Contains(mulContent, " ") {
					content = stripMulPrefix(content)
					continue
				}
				first, err := strconv.Atoi(values[0])
				if err != nil {
					content = stripMulPrefix(content)
					continue
				}
				second, err := strconv.Atoi(values[1])
				if err != nil {
					content = stripMulPrefix(content)
					continue
				}
				fmt.Println(content[startIndex : startIndex+len(mulContent)+5])
				out = append(out, []int{first, second})
				content = stripContentFromMulStatement(content, startIndex, mulContent)
			} else {
				content = stripMulPrefix(content)
			}
		} else {
			content = stripMulPrefix(content)
		}
	}
	return out
}

func sumMulStatements(statements [][]int) int {
	sum := 0
	for _, statement := range statements {
		sum += statement[0] * statement[1]
	}
	return sum
}

func stripContentFromMulStatement(content string, startIndex int, mulContent string) string {
	return content[startIndex+5+len(mulContent):]
}

func stripMulPrefix(content string) string {
	return content[strings.Index(content, "mul(")+4:]
}
