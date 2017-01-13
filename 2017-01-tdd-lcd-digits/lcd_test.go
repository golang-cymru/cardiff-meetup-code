package main

import "testing"

func Test_Numerals(t *testing.T) {

	lcd := map[int]string{
		0: "._.\n|.|\n|_|",
		1: "...\n..|\n..|",
		2: "._.\n._|\n|_.",
		3: "._.\n._|\n._|",
		4: "...\n|_|\n..|",
		5: "._.\n|_.\n._|",
		6: "._.\n|_.\n|_|",
		7: "._.\n..|\n..|",
		8: "._.\n|_|\n|_|",
		9: "._.\n|_|\n..|",
	}

	for k, v := range lcd {
		if got := printDigit(k); got != v {
			t.Errorf("Incorrect output: got '%v' expected '%v'", got, v)
		}
	}

}

func Test_Multiple_Digits(t *testing.T) {
	fortytwo := "... ._.\n|_| ._|\n..| |_."

	if got := printDigit(42); got != fortytwo {
		t.Errorf("Incorrect output: got '%v' expected '%v'", got, fortytwo)
	}
}
