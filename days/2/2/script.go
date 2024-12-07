package day02_01

import (
	"aoc2024/utils"
	"fmt"
	"math"
)

func Run() {
	lines := utils.GetAOCInputLines(2, 1)
	safeCount := 0
	for _, line := range lines {
		if isReportSave(utils.ParseLineToIntSlice(line, " ")) {
			safeCount++
		}
	}
	fmt.Println("Number of safe reports: ", safeCount)

}

func isReportSave(data []int) bool {
	unsafeCount := 0
	var isIncreasing bool
	for i, element := range data {
		isSave := true
		if i+1 >= len(data) {
			continue
		}
		if i == 0 {
			isIncreasing = element < data[i+1]
		}
		first, second := element, data[i+1]
		diff := first - second
		absDiff := math.Abs(float64(diff))
		isSave = isSave && absDiff >= 1 && absDiff <= 3
		isSave = isSave && ((isIncreasing && diff < 0) || (!isIncreasing && diff > 0))
		if !isSave {
			unsafeCount++
		}
	}
	return unsafeCount <= 1
}
