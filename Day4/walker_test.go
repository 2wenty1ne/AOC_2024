package main

import (
	"testing"
	"os"
	"strings"
)

func TestDirection(t *testing.T) {

	file, err := os.ReadFile("test.txt")

	if err != nil {
		panic("Error opening file, exeting!")
	}

	content := string(file)

	dirs := [...]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW", "N"}
	resultMap := map[string]string{
		"N": "OOOO",
		"NE": "PPPP",
		"E": "RRRR",
		"SE": "TTTT",
		"S": "UUUU",
		"SW": "KKKK",
		"W": "LLLL",
		"NW": "IIII"}

	lines := strings.Split(content, "\n")

	//? Clean whitespaces
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\r", "")
	}

	for _, dir := range dirs {
		walkResult := walkFullDir(lines, 4, 4, dir)

		expectedResult := resultMap[dir]
		if walkResult != expectedResult {
			t.Errorf("Result was incorrect, got: %s, want: %s.", walkResult, expectedResult)
		}
	}


}

