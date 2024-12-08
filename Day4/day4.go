package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	file, err := os.ReadFile("Data.txt")

	if err != nil {
		panic("Error opening file, exeting!")
	}

	content := string(file)

	lines := strings.Split(content, "\n")

	//? Clean whitespaces
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\r", "")
	}


	//? Star one
	totalHitSumStarOne := 0

	for y, line := range lines {
		for x, char := range line {
			str := string(char)

			if (str != "X") {continue}

			totalHitSumStarOne += processOneCharStarOne(lines, y, x)
		}
	}
	fmt.Printf("Total Hits first star: %v \n", totalHitSumStarOne)

	processOneCharStarTwo(lines, 3, 6)

	
	//? Star two
	totalHitSumStarTwo := 0

	for y, line := range lines {
		for x, char := range line {
			str := string(char)

			if (str != "A") {continue}

			totalHitSumStarTwo += processOneCharStarTwo(lines, y, x)
		}
	}
	fmt.Printf("Total Hits second star: %v \n", totalHitSumStarTwo)
}


func processOneCharStarTwo(lines []string, y int, x int) int {
	mCount := 0
	sCount := 0

	_ = mCount
	_ = sCount


	directions := [...]string{"NE", "SE", "SW", "NW"}

	for _, dir := range directions {
		if !checkDirection(lines, y, x, dir, 2) {continue}
		currrent, _, _ := walkDirection(lines, y, x, dir)
		if currrent == "S" {sCount++}
		if currrent == "M" {mCount++}
	}

	if sCount == 2 && mCount == 2 {
		return 1
	}

	return 0
}


func processOneCharStarOne(lines []string, y int, x int) int {
	directions := [...]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	hitSum := 0

	for _, dir := range directions {
		if !checkDirection(lines, y, x, dir, 4) {continue}
		hitSum += walkFullDir(lines, y, x, dir)
	}
	return hitSum
}


func walkFullDir(lines []string, y int, x int, dir string) int {
	needed := [...]string{"M", "A", "S"}

	for i := 0; i <=2; i++ {
		current, y_new, x_new := walkDirection(lines, y, x, dir)
		y = y_new
		x = x_new
		if current != needed[i] {return 0}
	}
	return 1
}


func checkDirection(lines []string, y int, x int, dir string, dist int) bool {
	switch dir {
	case "N":
		if !checkSizeNorth(y, dist) {return false}
	case "NE":
		if !checkSizeNE(lines, x, y, dist) {return false}
	case "E":
		if !checkSizeEast(lines, x, dist) {return false}
	case "SE":
		if !checkSizeSE(lines, x, y, dist) {return false}
	case "S":
		if !checkSizeSouth(lines, y, dist) {return false}
	case "SW":
		if !checkSizeSW(lines, x, y, dist) {return false}
	case "W":
		if !checkSizeWest(x, dist) {return false}
	case "NW":
		if !checkSizeNW(x, y, dist) {return false}
	default:
		panic("Unknown direction input into checkDirection")
	}
	return true
}


func walkDirection(lines []string, y int, x int, dir string) (string, int, int) {
	switch dir {
	case "N":
		y := y-1
		return string(lines[y][x]), y, x
	case "NE":
		y := y-1
		x := x+1
		return string(lines[y][x]), y, x
	case "E":
		x := x+1
		return string(lines[y][x]), y, x
	case "SE":
		y := y+1
		x := x+1
		return string(lines[y][x]), y, x
	case "S":
		y := y+1
		return string(lines[y][x]), y, x
	case "SW":
		y := y+1
		x := x-1
		return string(lines[y][x]), y, x
	case "W":
		x := x-1
		return string(lines[y][x]), y, x
	case "NW":
		y := y-1
		x := x-1
		return string(lines[y][x]), y, x
	}
	return "", 0, 0
}


//? Check directions
func checkSizeNorth(y int, dist int) bool {
	if y < dist-1 {return false} //? Was 3

	return true
}

func checkSizeEast(lines []string, x int, dist int) bool {
	remainingRight := len(lines[x]) - x
	if remainingRight < dist {return false} //? Was 4

	return true
}

func checkSizeSouth(lines []string, y int, dist int) bool {
	remainingDown := len(lines[y]) - y

	if remainingDown < dist {return false} //? Was 4

	return true
}

func checkSizeWest(x int, dist int) bool {
	if x < dist-1 {return false}

	return true
}

func checkSizeNE(lines []string, x int, y int, dist int) bool {
	return checkSizeNorth(y, dist) && checkSizeEast(lines, x, dist)
}

func checkSizeSE(lines []string, x int, y int, dist int) bool {
	return checkSizeEast(lines, x, dist) && checkSizeSouth(lines, y, dist)
}

func checkSizeSW(lines []string, x int, y int, dist int) bool {
	return checkSizeSouth(lines, y, dist) && checkSizeWest(x, dist)
}

func checkSizeNW(x int, y int, dist int) bool {
	return checkSizeWest(x, dist) && checkSizeNorth(y, dist)
}