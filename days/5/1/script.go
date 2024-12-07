package day05_01

import (
	"aoc2024/utils"
	"fmt"
)

func Run() {
	lines := utils.GetAOCInputLines(5, 1)
	// parse rules
	rules := parseRules(lines)
	instructions := parseInstructions(lines)
	correctInstructions := getCorrectInstructions(instructions, rules)
	for _, instinstruction := range correctInstructions {
		fmt.Println(instinstruction)
	}
	sum := 0
	for _, instruction := range correctInstructions {
		sum += getMiddlePageOfInstruction(instruction)
	}
	fmt.Println("Sum of middle pages: ", sum)

}

func parseRules(lines []string) [][]int {
	var out [][]int
	endOfRuleSection := false
	for _, line := range lines {
		if line == "" {
			endOfRuleSection = true
		}
		if endOfRuleSection {
			continue
		}
		out = append(out, utils.ParseLineToIntSlice(line, "|"))
	}
	return out
}

func parseInstructions(lines []string) [][]int {
	var out [][]int
	endOfRuleSection := false
	for _, line := range lines {
		if line == "" {
			endOfRuleSection = true
			continue
		}
		if !endOfRuleSection {
			continue
		}
		out = append(out, utils.ParseLineToIntSlice(line, ","))
	}
	return out
}

func getCorrectInstructions(instructions [][]int, rules [][]int) [][]int {
	var out [][]int
	for _, instruction := range instructions {
		fmt.Println(instruction)
		significantRules := getSignificantRulesForInstruction(instruction, rules)
		isValid := true
		for _, rule := range significantRules {
			if !isValid {
				continue
			}
			if utils.IndexOfArray(instruction, rule[0]) > utils.IndexOfArray(instruction, rule[1]) {
				isValid = false
			}
		}
		if isValid {
			out = append(out, instruction)
		}
	}
	fmt.Println("==========================")
	return out
}

func getSignificantRulesForInstruction(instruction []int, rules [][]int) [][]int {
	var out [][]int
	for _, rule := range rules {
		if utils.IndexOfArray(instruction, rule[0]) != -1 && utils.IndexOfArray(instruction, rule[1]) != -1 {
			out = append(out, rule)
		}
	}
	return out
}

func getMiddlePageOfInstruction(instruction []int) int {
	return instruction[len(instruction)/2]
}
