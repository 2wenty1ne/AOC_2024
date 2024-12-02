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

	for scanner.Scan() {
		report := scanner.Text();
		levels := strings.Split(report, " ");

		var ascending bool
		var unsafe bool
		
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
			continue;
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
				if !(element < compare) {unsafe = true}
			} else {
				if !(element > compare) {unsafe = true}
			}

			diff := abs(element - compare);
			if !(diff >= 1 && diff <= 3) {unsafe = true}
		}
		if !unsafe {
			amountSafeLevels = amountSafeLevels + 1;
		}
	}

	fmt.Printf("Result: %d \n", amountSafeLevels);

}


func abs(num int) int {
	if num < 0 {
		return -num;
	}
	return num;
}
