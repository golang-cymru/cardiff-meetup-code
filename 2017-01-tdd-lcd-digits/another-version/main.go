// Package main contains an example implementation of the LCD digit code exercise
// from Cyber Dojo (cyber-dojo.org)
//
// Taking a string of one or more digits it converts those into a string represantation
// of a 3x3 lcd digit grid (using pipes, dots and underscores)
//
// e.g.
//   $ go run main.go 42
//
// would be result in:
//
//   ... ._.
//   |_| ._|
//   ..| |_.
//
// being printed to STDOUT
//
// nb. This is presentated as example of code of one way the challenge may be
// tackled, it doesn't presume to be THE way or the best way to solve the problem.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Grab the number from the command line. First we'll make sure the user gave
	// us exactly one number to process.
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("Expected exactly one arg!")
		os.Exit(1)
	}
	number := arg[0]

	// Pass the number to the conversion routine and print the output
	fmt.Printf("\n%s\n", ConvertFromString(number))
}

// Define a mapping between a digit (rune) and its lcd repesenatation (three strings)
var lcd = map[rune][]string{
	'0': []string{"._.", "|.|", "|_|"},
	'1': []string{"...", "..|", "..|"},
	'2': []string{"._.", "._|", "|_."},
	'3': []string{"._.", "._|", "._|"},
	'4': []string{"...", "|_|", "..|"},
	'5': []string{"._.", "|_.", "._|"},
	'6': []string{"._.", "|_.", "|_|"},
	'7': []string{"._.", "..|", "..|"},
	'8': []string{"._.", "|_|", "|_|"},
	'9': []string{"._.", "|_|", "..|"},
}

// ConvertFromInt takes and integer and returns a string representation of the lcd display
func ConvertFromInt(number int) string {
	return ConvertFromString(strconv.Itoa(number))
}

// ConvertFromString takes a string and returns a string representation of the lcd display
func ConvertFromString(number string) string {

	// We want to store a set of blocks (items from the lcd map) that represent
	// the digits in our number. Once we have those we can read them out to get
	// our output.
	blocks := make([][]string, len(number), len(number))

	// Parse the digits in the input and get the matching block
	// for each. To do that we can just range over the strig to get
	// each rune in turn.
	for i, d := range number {
		blocks[i] = lcd[d]
	}

	// Initialise an empty string to
	output := ""

	// Now we have a set of digits, we need to format them row by row
	for row := 0; row < 3; row++ {
		// Loop through this row of each digit to get the full line
		for _, block := range blocks {
			output += block[row] + " "
		}
		// Remove the extraneous extra space that we've got after the last one
		output = strings.TrimSuffix(output, " ") + "\n"

	}
	return output
}
