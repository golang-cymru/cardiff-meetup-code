package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1:][0]
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("That is not a number!")
		os.Exit(1)
	}

	out := printDigit(num)
	fmt.Println(out)

}

type lcdDigit struct {
	matrix [3][3]string
}

type lcdDigits struct {
	digitsMap map[int]lcdDigit
}

func (lD lcdDigits) makeLCDScreen(num int) string {
	//below will not work with single digit zero so return result for that
	if num == 0 {
		return lD.digitsMap[0].matrix[0][0] +
			lD.digitsMap[0].matrix[0][1] +
			lD.digitsMap[0].matrix[0][2] + "\n" +
			lD.digitsMap[0].matrix[1][0] +
			lD.digitsMap[0].matrix[1][1] +
			lD.digitsMap[0].matrix[1][2] + "\n" +
			lD.digitsMap[0].matrix[2][0] +
			lD.digitsMap[0].matrix[2][1] +
			lD.digitsMap[0].matrix[2][2]
	}
	i := uint64(num)
	start := i
	b64 := uint64(10)
	var lines [3]string
	for ; i > 0; i /= b64 {
		rightMostDigit := int(i % b64)
		for i2 := 0; i2 < 3; i2++ {
			if i == start && i2 != 2 {
				lines[i2] = "\n"
			} else {
				lines[i2] = " " + lines[i2]
			}
			lines[i2] = lD.digitsMap[rightMostDigit].matrix[i2][0] +
				lD.digitsMap[rightMostDigit].matrix[i2][1] +
				lD.digitsMap[rightMostDigit].matrix[i2][2] +
				lines[i2]
		}
	}
	return lines[0] + lines[1] + strings.TrimSuffix(lines[2], " ")
}

var digits = lcdDigits{
	digitsMap: map[int]lcdDigit{
		0: lcdDigit{[3][3]string{{".", "_", "."}, {"|", ".", "|"}, {"|", "_", "|"}}},
		1: lcdDigit{[3][3]string{{".", ".", "."}, {".", ".", "|"}, {".", ".", "|"}}},
		2: lcdDigit{[3][3]string{{".", "_", "."}, {".", "_", "|"}, {"|", "_", "."}}},
		3: lcdDigit{[3][3]string{{".", "_", "."}, {".", "_", "|"}, {".", "_", "|"}}},
		4: lcdDigit{[3][3]string{{".", ".", "."}, {"|", "_", "|"}, {".", ".", "|"}}},
		5: lcdDigit{[3][3]string{{".", "_", "."}, {"|", "_", "."}, {".", "_", "|"}}},
		6: lcdDigit{[3][3]string{{".", "_", "."}, {"|", "_", "."}, {"|", "_", "|"}}},
		7: lcdDigit{[3][3]string{{".", "_", "."}, {".", ".", "|"}, {".", ".", "|"}}},
		8: lcdDigit{[3][3]string{{".", "_", "."}, {"|", "_", "|"}, {"|", "_", "|"}}},
		9: lcdDigit{[3][3]string{{".", "_", "."}, {"|", "_", "|"}, {".", ".", "|"}}},
	},
}

func printDigit(number int) string {
	return digits.makeLCDScreen(number)
}
