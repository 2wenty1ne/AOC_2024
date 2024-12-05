package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	file, err := os.ReadFile("testData.txt")

	if err != nil {
		panic("Error opening file, exeting!")
	}

	content := string(file)

	lines := strings.Split(content, "\n")

	//? Clean whitespaces
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\r", "")
	}
	//mid_y, mid_x := 4
	directions := [...]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW", "N"}
	_ = directions
	
	test_line := lines[4]



	for t := 4; t < len(test_line)-1; t++ {
		fmt.Printf("%v -> %v \n", string(lines[t][t]), walkNW(lines, t, t))
	}


// 	for y, line := range lines {
// 		fmt.Printf("Y: %v\n", y)
// 		for x, char := range line {
// 			str := string(char)
// 			if (str != "M") {continue}

// 			sizeLeft := checkSizeWest(x)
// 			sizeRight := checkSizeEast(lines, x)
// 			fmt.Printf("Hor: %v - %v - %v \n", sizeLeft, x, sizeRight)

// 			sizeUp := checkSizeNorth(y)
// 			sizeDown := checkSizeSouth(lines, y)
// 			fmt.Printf("Vert: %v - %v - %v \n", sizeUp, y, sizeDown)

// 			fmt.Println()
// 		}
// 	}
}

func walkFullDir(lines []string, y_start int, x_start int, dir string) string {
	possible := checkDirection(lines, y_start, x_start, dir)

	if !possible {return ""}

	wholeString := ""
	for i := 0; i <=3; i++ {
		wholeString += walkDirection(lines, y_start, x_start, dir)
	}

	return wholeString
}

func checkDirection(lines []string, y_start int, x_start int, dir string) bool {
	switch dir {
	case "N":
		if !checkSizeNorth(y_start) {return false}
	case "NE":
		if !checkSizeNE(lines, x_start, y_start) {return false}
	case "E":
		if !checkSizeEast(lines, x_start) {return false}
	case "SE":
		if !checkSizeSE(lines, x_start, y_start) {return false}
	case "S":
		if !checkSizeSouth(lines, y_start) {return false}
	case "SW":
		if !checkSizeSW(lines, x_start, y_start) {return false}
	case "W":
		if !checkSizeWest(x_start) {return false}
	case "NW":
		if !checkSizeNW(x_start, y_start) {return false}
	}
	return true
}

func walkDirection(lines []string, y_start int, x_start int, dir string) string {
	switch dir {
	case "N":
		return walkNorth(lines, y_start, x_start)
	case "NE":
		return walkNE(lines, y_start, x_start)
	case "E":
		return walkEast(lines, y_start, x_start)
	case "SE":
		return walkSE(lines, y_start, x_start)
	case "S":
		return walkSouth(lines, y_start, x_start)
	case "SW":
		return walkSW(lines, y_start, x_start)
	case "W":
		return walkWest(lines, y_start, x_start)
	case "NW":
		return walkNW(lines, y_start, x_start)
	}
	return ""
}

//? Next char in direction
func walkNorth(lines []string, y int, x int) string {
	return string(lines[y-1][x])
}

func walkNE(lines []string, y int, x int) string {
	return string(lines[y-1][x+1])
}

//Todo test
func walkEast(lines []string, y int, x int) string {
	return string(lines[y][x+1])
}

func walkSE(lines []string, y int, x int) string {
	return string(lines[y+1][x+1])
}

func walkSouth(lines []string, y int, x int) string {
	return string(lines[y+1][x])
}

func walkSW(lines []string, y int, x int) string {
	return string(lines[y+1][x-1])
}

//TODO test
func walkWest(lines []string, y int, x int) string {
	return string(lines[y][x-1])
}


func walkNW(lines []string, y int, x int) string {
	return string(lines[y-1][x-1])
}


//? Check directions
func checkSizeNorth(y int) bool {
	if y < 3 {return false}

	return true
}

func checkSizeEast(lines []string, x int) bool {
	remainingRight := len(lines[x]) - x //? 10 - 1

	if remainingRight < 4 {return false}

	return true
}

func checkSizeSouth(lines []string, y int) bool {
	remainingDown := len(lines[y]) - y

	if remainingDown < 4 {return false}

	return true
}

func checkSizeWest(x int) bool {
	if x < 3 {return false}

	return true
}

func checkSizeNE(lines []string, x int, y int) bool {
	return checkSizeNorth(y) && checkSizeEast(lines, x)
}

func checkSizeSE(lines []string, x int, y int) bool {
	return checkSizeEast(lines, x) && checkSizeSouth(lines, y)
}

func checkSizeSW(lines []string, x int, y int) bool {
	return checkSizeSouth(lines, y) && checkSizeWest(x)
}

func checkSizeNW(x int, y int) bool {
	return checkSizeWest(x) && checkSizeNorth(y)
}