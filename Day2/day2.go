package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	//"reflect"
)



func main() {
	file, err := os.Open("Data.txt");

	if err != nil {
		fmt.Println("Opening File failed!");
		return;
	}
	defer file.Close();

	amountSafeLevels := 0;

	scanner := bufio.NewScanner(file);

	//? Go through file line by line
	for scanner.Scan() {
		report := scanner.Text();
		levels := strings.Split(report, " ");

		var safe bool

		for i := 0; i < len(levels); i++ {
			copyLevels := append([]string{}, levels...)
			copyLevels = append(copyLevels[:i], copyLevels[i+1:]...)
			unsafeCounter := checkReport(copyLevels)

			if unsafeCounter <= 1 {safe = true}
		}

		if safe {amountSafeLevels++}

		// if !unsafe {
		// 	amountSafeLevels = amountSafeLevels + 1;
		// }
	}

	fmt.Printf("Result: %d \n", amountSafeLevels);

}


func removeByIndex(arr []string, index int) []string {
	return append(arr[:index], arr[index+1:]...)
}


func checkReport(levels []string) int {
	var ascending bool
	counter := 0 

	//? Check the first two to determine ascending / descending
	first, err := strconv.Atoi(levels[0]);
		if err != nil {
			panic(err);
		}
	second, err := strconv.Atoi(levels[1]);
		if err != nil {
			panic(err);
		}
	
	//? Determine ascending / descending
	if (first < second) {
		ascending = true;
	}

	//? Check if first two are not meeting difference requirement
	firstDiff := abs(first - second);
	if !(firstDiff >= 1 && firstDiff <= 3) {
		counter ++
	}

	//? Check remaining levels in reports
	for i := 1; i < (len(levels) - 1); i++ {
		element, err := strconv.Atoi(levels[i]);
		if err != nil {
			panic(err);
		}

		compare, err := strconv.Atoi(levels[i+1]);
		if err != nil {
			panic(err);
		}

		if ascending {
			if !(element < compare) {counter ++}
		} else {
			if !(element > compare) {counter ++}
		}

		diff := abs(element - compare);
		if !(diff >= 1 && diff <= 3) {counter ++}
	}

	return counter
}


func abs(num int) int {
	if num < 0 {
		return -num;
	}
	return num;
}
