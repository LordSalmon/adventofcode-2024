package day02_01

import (
	"aoc2024/utils"
)

func Run() {
	lines := utils.GetAOCInputLines(2, 1)
	var data []int
	for _, line := range lines {
		data = append(data, utils.ParseLineToIntSlice(line)...)
	}

}

func isReportSave(data []int) bool {
	isSave := true
	for i, element := range data {
		if i+2 >= len(data) {
			continue
		}
		isSave = isSave || isDataDifferenceSave(element, data[i+1])
	}
	return isSave
}

func isDataDifferenceSave(first int, second int) bool {
	diff := first - second
	return diff >= -1 && diff <= 3
}
