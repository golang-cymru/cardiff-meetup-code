package main

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type stringTestCase struct {
	Number   string
	Expected string
}

func TestConvert(t *testing.T) {

	// Define a test matrix of inputs to expected outputs
	tests := []stringTestCase{
		stringTestCase{Number: "1", Expected: "...\n..|\n..|\n"},
		stringTestCase{Number: "42", Expected: "... ._.\n|_| ._|\n..| |_.\n"},
	}

	// Using go convey for some nicer test furniture - nothing you can't
	// logically do with the standard testing package.
	Convey("When converting from a string", t, func() {
		for _, test := range tests {
			Convey(fmt.Sprintf("Given an input of '%s'", test.Number), func() {
				So(ConvertFromString(test.Number), ShouldEqual, test.Expected)
			})
		}
	})

	// To test both our interface methods we can just reuse our existing test
	// matrix and just convert the inputs to integers before feeding them in.
	Convey("When converting from an integer", t, func() {
		for _, test := range tests {
			// For brevity we're throwing away the error here. Kinda acceptable in this
			// test as we have full control over the data to ensure it's correctly
			// convertable and error checking would make the code flow more complex
			// for little gain, but as a rule you shouldn't ignore errors!
			number, _ := strconv.Atoi(test.Number)
			Convey(fmt.Sprintf("Given an input of '%d'", number), func() {
				So(ConvertFromInt(number), ShouldEqual, test.Expected)
			})
		}
	})

}
