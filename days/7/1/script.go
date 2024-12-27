package day07_01

import (
	"aoc2024/utils"
	"fmt"
	"math"
)

func Run() {
	lines := utils.GetAOCInputLines(7, 1)
	lines = lines
	sum := 0
	println(canEquationEqualResult(190, []int{10, 19}))
	println(canEquationEqualResult(3267, []int{81, 40, 27}))
	println(canEquationEqualResult(83, []int{17, 5}))
	println(canEquationEqualResult(156, []int{15, 6}))
	println(canEquationEqualResult(7290, []int{6, 8, 6, 15}))
	println(canEquationEqualResult(292, []int{11, 6, 16, 20}))
	// for _, line := range lines {
	// 	result := utils.MustAtoi(strings.Split(line, ":")[0])
	// 	equationPars := utils.ParseLineToIntSlice(strings.Split(line, ": ")[1], " ")
	// 	if canEquationEqualResult(result, equationPars) {
	// 		sum += result
	// 	}
	// }

	fmt.Println("Sum of correct equations: ", sum)
}

func canEquationEqualResult(result int, parts []int) bool {
	println("testing for ", result)
	for operator := 0; operator < int(math.Pow(float64(2), float64(len(parts)-1))); operator++ {
		println("operator: ", operator)
		stack := []int{parts[0]}
		for i := 1; i < len(parts); i++ {
			if operator&(int(math.Pow(float64(2), float64(i-1)))) > 0 {
				println("adding to stack", parts[i])
				stack = append(stack, parts[i])
			} else {
				println("multiplying last element in stack with", parts[i])
				lastStackElement := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, lastStackElement*parts[i])
			}
		}
		currentResult := 0
		for _, part := range stack {
			currentResult += part
		}
		if currentResult == result {
			return true
		}
	}
	return false
}
