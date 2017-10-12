package testing

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct{
		name string
		input string
		expectedOutput string
	}{
		{
			name:"empty string",
			input: "",
			expectedOutput: "",
		},
		{
			name:"single character",
			input: "a",
			expectedOutput: "a",
		}, 
		{
			name:"two characters",
			input: "ab",
			expectedOutput: "ba",
		},
		{
			name:"palindrome",
			input: "stressed",
			expectedOutput: "desserts",
		},
	}

	for _, c := range cases {
		if output := Reverse(c.input); output != c.expectedOutput {
			t.Errorf("%s: got %s but expected %s", c.name, c.input, c.expectedOutput)
		}
	}
}

func TestGetGreeting(t *testing.T) {
	cases := []struct{
		name string
		input string
		expectedOutput string
	} {
		{
		name:"empty string",
		input: "",
		expectedOutput: "Hello World!",
		}, 
		{
			name:"non-empty string",
			input: "Billy",
			expectedOutput: "Hello ",
		},
	}
	for _, c := range cases {
		if output := GetGreeting(c.input); output != c.expectedOutput {
			t.Errorf("%s: got %s but expected %s", c.name, c.input, c.expectedOutput)
		}
	}
}

func TestParseSize(t *testing.T) {
	cases := []struct{
		name string
		input string
		expectedOutput *Size
	} {
		{
		name:"empty string",
		input: "",
		expectedOutput: &Size{},
		}, 
	}
	for _, c := range cases {
		if output := ParseSize(c.input); output.Height != c.expectedOutput.Height || output.Width != c.expectedOutput.Width {
			t.Errorf("%v: got %v but expected %v", c.name, c.input, c.expectedOutput)
		}
	}
}

func TestLateDaysConsume(t *testing.T) {
	ld := NewLateDays()
	for i := 3; i > -10; i-- {
		expectedLateDays := i
		if expectedLateDays < 0 {
			expectedLateDays = 0
		}
		if output := ld.Consume("test"); output != expectedLateDays {
			t.Errorf("iteration %d: got %d but expected %d", i, output, expectedLateDays)
		}
	}
}
