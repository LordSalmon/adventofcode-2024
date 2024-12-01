package day01_01

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Run() {
	lines := utils.GetAOCInputLines(1, 1)
	distance := 0
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
	for len(firstList) > 0 {
		smallestFirst := getSmallesElementFromList(firstList)
		smallestSecond := getSmallesElementFromList(secondList)
		distance += int(math.Abs(float64(smallestSecond - smallestFirst)))
		firstList = removeElementFromList(firstList, smallestFirst)
		secondList = removeElementFromList(secondList, smallestSecond)
	}
	fmt.Println("Distance: ", distance)
}

func getSmallesElementFromList(list []int) int {
	smallest := list[0]
	for _, element := range list {
		if element < smallest {
			smallest = element
		}
	}
	return smallest
}

func removeElementFromList(list []int, element int) []int {
	var newList []int
	alreadyAvoided := false
	for _, listElement := range list {
		if listElement != element || alreadyAvoided {
			newList = append(newList, listElement)
		} else {
			alreadyAvoided = true
		}
	}
	return newList
}
