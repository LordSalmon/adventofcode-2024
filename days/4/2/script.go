package day04_02

import (
	"aoc2024/utils"
	"fmt"
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
	for i := range lines {
		if i == 0 || i >= len(lines)-1 {
			continue
		}
		for j := range lines[i] {
			if j == 0 || j >= len(lines[i])-1 {
				continue
			}
			if string(lines[i][j]) == "A" {
				if string(lines[i-1][j-1]) == "M" && string(lines[i+1][j+1]) == "S" || string(lines[i-1][j-1]) == "S" && string(lines[i+1][j+1]) == "M" {
					if string(lines[i-1][j+1]) == "M" && string(lines[i+1][j-1]) == "S" || string(lines[i-1][j+1]) == "S" && string(lines[i+1][j-1]) == "M" {
						count++
					}
				}
			}
		}
	}
	fmt.Println("Count: ", count)
}
