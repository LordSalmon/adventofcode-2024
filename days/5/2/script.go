package day05_02

import (
	"aoc2024/utils"
	"fmt"
)

func Run() {
	lines := utils.GetAOCInputLines(5, 1)
	// parse rules
	rules := parseRules(lines)
	instructions := parseInstructions(lines)
	instructions = getInorrectInstructions(instructions, rules)
	instructions = correctInstructions(instructions, rules)
	sum := 0
	for _, instruction := range instructions {
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

func getInorrectInstructions(instructions [][]int, rules [][]int) [][]int {
	var out [][]int
	for _, instruction := range instructions {
		significantRules := getSignificantRulesForInstruction(instruction, rules)
		isValid := true
		for _, rule := range significantRules {
			if !isValid {
				continue
			}
			if !isRuleValid(instruction, rule) {
				isValid = false
			}
		}
		if !isValid {
			out = append(out, instruction)
		}
	}
	return out
}

func isRuleValid(instruction []int, rule []int) bool {
	return utils.IndexOfArray(instruction, rule[0]) < utils.IndexOfArray(instruction, rule[1])
}

func getViolatingRules(instructions []int, rules [][]int) [][]int {
	var out [][]int
	for _, rule := range rules {
		if !isRuleValid(instructions, rule) {
			out = append(out, rule)
		}
	}
	return out
}

func correctInstructions(instructions [][]int, rules [][]int) [][]int {
	var out [][]int
	for _, instruction := range instructions {
		significantRules := getSignificantRulesForInstruction(instruction, rules)
		violatingRules := getViolatingRules(instruction, significantRules)
		newInstruction := []int{}
		for len(violatingRules) > 0 {
			// switch numbers of first violating rule in instruction
			rule := violatingRules[0]
			newInstruction = append(newInstruction, instruction[:utils.IndexOfArray(instruction, rule[1])]...)
			newInstruction = append(newInstruction, rule[0])
			newInstruction = append(newInstruction, instruction[utils.IndexOfArray(instruction, rule[1])+1:utils.IndexOfArray(instruction, rule[0])]...)
			newInstruction = append(newInstruction, rule[1])
			newInstruction = append(newInstruction, instruction[utils.IndexOfArray(instruction, rule[0])+1:]...)
			instruction = newInstruction
			newInstruction = []int{}
			violatingRules = getViolatingRules(instruction, significantRules)
		}
		out = append(out, instruction)
	}
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
