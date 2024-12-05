package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.ReadFile("Data.txt")

	if err != nil {
		panic("Error opening file, exeting!")
	}

	content := string(file)

	sum := 0

	for i, char := range content {
		remainingChars := len(content) - i
		firstWholeNumStr := ""

		if char != 'm' {
			continue
		}

		if remainingChars < 8 { break }


		//? Check if next 3 are ul(
		nextThree := string(content[i+1]) + string(content[i+2]) + string(content[i+3]) 

		if nextThree != "ul(" { continue }


		//? Check if first element in () is a number
		_, errFirstNum := strconv.Atoi(string(content[i+4]))

		if errFirstNum != nil { continue }

		firstWholeNumStr += string(content[i+4])


		//? Check if second element in () is number or komma
		_, errSecondElem := strconv.Atoi(string(content[i+5]))
		if errSecondElem != nil {
			if string(content[i+5]) != "," {
				continue
			}

			product, continueErr := computeFirstValue(i, content, i+5+1, firstWholeNumStr)
			if continueErr != nil {continue}
			sum += product
			continue
		}
		firstWholeNumStr += string(content[i+5])


		//? Check if third element in () is number or komma
		_, errThirdElem := strconv.Atoi(string(content[i+6]))
		if errThirdElem != nil {
			if string(content[i+6]) != "," {
				continue
			}
			product, continueErr := computeFirstValue(i, content, i+6+1, firstWholeNumStr)
			if continueErr != nil {
				continue
			}
			sum += product
			continue
		}
		firstWholeNumStr += string(content[i+6])


		if string(content[i+7]) != "," {continue}
		product, continueErr := computeFirstValue(i, content, i+7+1, firstWholeNumStr)
		if continueErr != nil {
			if continueErr.Error() == "continue" {continue}
			if continueErr.Error() == "break" {break}
		}
		sum += product
	}
	fmt.Printf("Result: %v \n", sum)
}


func computeFirstValue(i int, content string, offset int, firstNumStr string) (int, error)  {
	firstNum, err := strconv.Atoi(firstNumStr)
	if (err != nil) {
		panic(fmt.Sprintf("Error converting first str to num: |%v| Type: %T \n", firstNumStr, firstNumStr))
	}
	
	secondNum, continueErr := checkAfterKomma(i, content, offset)

	return firstNum * secondNum, continueErr
}


func checkAfterKomma(i int, content string, offset int) (int, error) {
	remainingChars := len(content) - i
	secondWholeNumStr := ""
	offsetComp := 8 + (offset - i) - 6

	//? Check if theres space for minimum
	if (offset - i) > 6 {
		// fmt.Printf("Break offset: %v \nReal offset: %v\nRemaining: %v \n", offset, offset-i,remainingChars) //! TEST
		// fmt.Printf("OffsetComp: %v \n", offsetComp) //!TEST
		if remainingChars < offsetComp { return 0, fmt.Errorf("break") } 
	}

	// fmt.Printf("|%v|, |%v|, |%v| \n", string(content[offset]), string(content[offset+1]), string(content[offset+2]))

	//? Check if first element is int
	_, errFirstNum := strconv.Atoi(string(content[offset]))

	if errFirstNum != nil { return 0, fmt.Errorf("continue") }

	secondWholeNumStr += string(content[offset])


	//? Check second element after komma
	_, errSecElem := strconv.Atoi(string(content[offset+1]))
	if errSecElem != nil {
		if string(content[offset+1]) != ")" {
			return 0, fmt.Errorf("continue")
		}
		firstNum, err := strconv.Atoi(secondWholeNumStr)
		if err != nil {panic(fmt.Sprintf("Error converting second str to num: %v \n", secondWholeNumStr))} 

		return firstNum, nil
	}
	secondWholeNumStr += string(content[offset+1])


	//? Check if there is space for third element after komma
	// fmt.Printf("Break offset: %v \nReal offset: %v\nRemaining: %v \n", offset, offset-i,remainingChars) //! TEST
	// fmt.Printf("OffsetComp: %v \n", offsetComp) //!TEST
	if remainingChars < offsetComp + 1 { return 0, fmt.Errorf("break") }

	//? Check third element after komma
	_, errThirdElem := strconv.Atoi(string(content[offset+2]))
	if errThirdElem != nil {
		if string(content[offset+2]) != ")" {
			return 0, fmt.Errorf("continue")
		}
		firstNum, err := strconv.Atoi(secondWholeNumStr)
		if err != nil {panic(fmt.Sprintf("Error converting second str to num: %v \n", secondWholeNumStr))} 

		return firstNum, nil
	}
	secondWholeNumStr += string(content[offset+2])


	//? Check if there is space for last ( after komma
	// fmt.Printf("Break offset: %v \nReal offset: %v\nRemaining: %v \n", offset, offset-i,remainingChars) //! TEST
	// fmt.Printf("OffsetComp: %v \n", offsetComp) //!TEST
	if remainingChars < offsetComp + 2 { return 0, fmt.Errorf("break") }

	if string(content[offset+3]) != ")" {
		return 0, fmt.Errorf("continue")
	}
	firstNum, err := strconv.Atoi(secondWholeNumStr)
	if err != nil {panic(fmt.Sprintf("Error converting second str to num: %v \n", secondWholeNumStr))} 

	return firstNum, nil
}
