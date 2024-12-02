package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadAOCInput(day int, part int) string {
	data, err := os.ReadFile("./inputs/day" + strconv.Itoa(day) + "_" + strconv.Itoa(part) + ".txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func GetAOCInputLines(day int, part int) []string {
	return strings.Split(ReadAOCInput(day, part), "\n")
}

func ParseLineToIntSlice(line string) []int {
	var data []int
	for _, element := range strings.Split(line, " ") {
		value, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}
		data = append(data, value)
	}
	return data
}
