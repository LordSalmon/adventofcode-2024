package main

import (
	"os"
)

func readAOCInput(day int, part int) string {
	file, err := os.Open("./inputs/day" + string(day) + "_" + string(part) + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input string
	_, err = file.Read([]byte(input))
	if err != nil {
		panic(err)
	}

	return input
}
