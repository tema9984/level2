package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		str       string
		ind       []int
		delim     string
		expected  string
		separated bool
	}{
		{
			str: "one	two	three	first	second	third",
			ind:   []int{1, 2},
			delim: "\t",
			expected: "two	three",
			separated: true,
		},
		{
			str:       "one two three first second third",
			ind:       []int{1, 2},
			delim:     "\t",
			expected:  "",
			separated: true,
		},
		{
			str:       "one-two-three-first-second-third",
			ind:       []int{1, 2},
			delim:     "-",
			expected:  "two-three",
			separated: true,
		},
		{
			str:       "one two three first second third",
			ind:       []int{0},
			delim:     "-",
			expected:  "one two three first second third",
			separated: false,
		},
	}
	//Default
	for _, tC := range testTable {
		result := serparat(tC.str, tC.ind, tC.delim, tC.separated)

		t.Logf("Calling serparat(%s), result %s\n", tC.str, result)
		//Assert
		if tC.expected != result {
			t.Errorf("Incorrect result, Exept %s, got %s", tC.expected, result)
		}

	}

}
