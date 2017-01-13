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

var numbers = map[string][]string{
	"0": []string{"._.", "|.|", "|_|"},
	"1": []string{"...", "..|", "..|"},
	"2": []string{"._.", "._|", "|_."},
	"3": []string{"._.", "._|", "._|"},
	"4": []string{"...", "|_|", "..|"},
	"5": []string{"._.", "|_.", "._|"},
	"6": []string{"._.", "|_.", "|_|"},
	"7": []string{"._.", "..|", "..|"},
	"8": []string{"._.", "|_|", "|_|"},
	"9": []string{"._.", "|_|", "..|"},
}

func printDigit(number int) string {
	str := strconv.Itoa(number)

	line1, line2, line3 := "", "", ""
	for i := 0; i < len(str); i++ {
		line1 += numbers[string(str[i])][0] + " "
		line2 += numbers[string(str[i])][1] + " "
		line3 += numbers[string(str[i])][2] + " "
	}
	return strings.TrimSuffix(line1, " ") + "\n" + strings.TrimSuffix(line2, " ") + "\n" + strings.TrimSuffix(line3, " ")
}
