package day06_01

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func Run() {
	content := utils.GetAOCInputLines(6, 1)
	roomWidth := len(content[0])
	roomHeight := len(content)
	currentPosition := getStartingPositionOfGuard(content)
	visitedLocations := []Position{currentPosition}
	direction := "up"

	for !isNextStepOufOfRoom(currentPosition, direction, roomWidth, roomHeight) {
		if isObstacleInFrontOfGuard(content, currentPosition, direction) {
			direction = turnDirection(direction)
		}
		currentPosition = getNextPosition(currentPosition, direction)
		visitedLocations = addVisitedLocation(currentPosition, visitedLocations)
	}
	fmt.Println("Visited locations: ", len(visitedLocations))
}

func addVisitedLocation(location Position, visitedLocations []Position) []Position {
	for _, visitedLocation := range visitedLocations {
		if visitedLocation == location {
			return visitedLocations
		}
	}
	return append(visitedLocations, location)
}

type Position struct {
	X int
	Y int
}

func getStartingPositionOfGuard(content []string) Position {
	for y, line := range content {
		if strings.Contains(line, "^") {
			position := Position{}
			position.X = strings.Index(line, "^")
			position.Y = y
			return position
		}
	}
	return Position{}
}

func getNextPosition(position Position, direction string) Position {
	switch direction {
	case "up":
		position.Y--
	case "down":
		position.Y++
	case "left":
		position.X--
	case "right":
		position.X++
	}
	return position
}

func isNextStepOufOfRoom(position Position, direction string, roomWidth int, roomHeight int) bool {
	nextStep := getNextPosition(position, direction)
	if nextStep.X < 0 || nextStep.X >= roomWidth {
		return true
	}
	if nextStep.Y < 0 || nextStep.Y >= roomHeight {
		return true
	}
	return false
}

func turnDirection(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	}
	return "up"
}

func isObstacleInFrontOfGuard(content []string, position Position, direction string) bool {
	switch direction {
	case "up":
		return content[position.Y-1][position.X] == '#'
	case "down":
		return content[position.Y+1][position.X] == '#'
	case "left":
		return content[position.Y][position.X-1] == '#'
	case "right":
		return content[position.Y][position.X+1] == '#'
	}
	return false
}
