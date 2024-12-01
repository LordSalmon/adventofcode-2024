package day01_02

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	lines := utils.GetAOCInputLines(1, 1)
	similarity := 0
	var firstList []int
	var secondList []int
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		firstValue, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			panic(err)
		}
		secondValue, err := strconv.Atoi(lineSplit[len(lineSplit)-1])
		if err != nil {
			panic(err)
		}
		firstList = append(firstList, firstValue)
		secondList = append(secondList, secondValue)
	}
	for _, number := range firstList {
		similarity += number * countElementsInList(secondList, number)
	}
	fmt.Println("Similarity: ", similarity)
}

func countElementsInList(list []int, element int) int {
	count := 0
	for _, listElement := range list {
		if listElement == element {
			count++
		}
	}
	return count
}
