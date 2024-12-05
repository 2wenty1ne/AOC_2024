package main

import (
	_ "fmt"
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

	directions := [...]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW", "N"}
	_ = directions

	for y, line := range lines {
		for x, char := range line {
			str := string(char)
			_ = x
			_ = y
			if (str != "X") {continue}
		}
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

func processOneChar(lines []string, y int, x int) int {
	directions := [...]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW", "N"}
	hitSum := 0

	for _, dir := range directions {
		if !checkDirection(lines, y, x, dir) {continue}


	}

	return hitSum
}

func walkFullDir(lines []string, y_start int, x_start int, dir string) int {
	isPossible := checkDirection(lines, y_start, x_start, dir)

	if !isPossible {return 0}

	needed := [...]string{"M", "A", "S"}

	for i := 0; i <=3; i++ {
		current := walkDirection(lines, y_start, x_start, dir)
		if current != needed[i] {return 0}
	}

	return 1
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
		return string(lines[y_start-1][x_start])
	case "NE":
		return string(lines[y_start-1][x_start+1])
	case "E":
		return string(lines[y_start][x_start+1])
	case "SE":
		return string(lines[y_start+1][x_start+1])
	case "S":
		return string(lines[y_start+1][x_start])
	case "SW":
		return string(lines[y_start+1][x_start-1])
	case "W":
		return string(lines[y_start][x_start-1])
	case "NW":
		return string(lines[y_start-1][x_start-1])
	}
	return ""
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