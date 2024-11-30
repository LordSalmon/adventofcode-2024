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
