package day04_02

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func Run() {
	//content := utils.ReadAOCInput(3, 1)
	lines := utils.GetAOCInputLines(4, 1)
	/**
	 * GET
	 * line forwards
	 * line backwards
	 * vertical down
	 * verticalup
	 * diagonal to left down
	 * diagonal to left up
	 * diagonal to right down
	 * diagonal to right up
	 */
	count := 0
	for _, line := range lines {
		count += getNumberOfOccurrencesInLine(line)
		count += getNumberOfOccurrencesInLine(reverseString(line))
	}
	for i := 0; i < len(lines[0]); i++ {
		topDownLine := getTopdownLine(lines, i)
		count += getNumberOfOccurrencesInLine(topDownLine)
		count += getNumberOfOccurrencesInLine(reverseString(topDownLine))
	}
	tldrLines := getDiagonalLines(lines)
	for _, line := range tldrLines {
		count += getNumberOfOccurrencesInLine(line)
		count += getNumberOfOccurrencesInLine(reverseString(line))
	}
	dltrLines := getDiagonalLines(rotateMatrix(lines))
	for _, line := range dltrLines {
		count += getNumberOfOccurrencesInLine(line)
		count += getNumberOfOccurrencesInLine(reverseString(line))
	}
	fmt.Println("Count: ", count)
}

func getNumberOfOccurrencesInLine(line string) int {
	occurrences := 0
	for strings.Contains(line, "XMAS") {
		occurrences++
		line = line[strings.Index(line, "XMAS")+1:]
	}
	return occurrences
}

func reverseString(s string) string {
	out := ""
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}
	return out
}

func getTopdownLine(lines []string, index int) string {
	out := ""
	for _, line := range lines {
		out += string(line[index])
	}
	return out
}

func getDiagonalLines(lines []string) []string {
	var out []string
	for i := range lines[0] {
		diagonalLine := ""
		for j, line := range lines {
			if len(line) >= i+j+1 {
				diagonalLine += string(line[i+j])
			}
		}
		out = append(out, diagonalLine)
	}
	topdownLine := getTopdownLine(lines, 0)
	for i := range topdownLine {
		if i == 0 {
			continue
		}
		diagonalLine := ""
		for j, line := range lines {
			if j-i >= 0 {
				diagonalLine += string(line[j-i])
			}
		}
		out = append(out, diagonalLine)
	}
	return out
}

func rotateMatrix(lines []string) []string {
	var out []string
	for i := range lines[0] {
		out = append(out, reverseString(getTopdownLine(lines, i)))
	}
	return out
}

func reverseLines(lines []string) []string {
	var out []string
	for _, line := range lines {
		out = append(out, reverseString(line))
	}
	return out
}
