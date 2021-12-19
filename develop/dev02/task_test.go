package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		str      string
		expected string
	}{
		{
			str:      "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			str:      "abcd",
			expected: "abcd",
		},
		{
			str:      "a4bc2d5e5",
			expected: "aaaabccdddddeeeee",
		},
		{
			str:      "",
			expected: "",
		},
		{
			str:      `qwe\4\5`,
			expected: "qwe45",
		},
		{
			str:      `qwe\45`,
			expected: "qwe44444",
		},
		{
			str:      `qwe\\5`,
			expected: `qwe\\\\\`,
		},
	}
	//Act
	for _, testCase := range testTable {
		result, err := unpacking(testCase.str)
		t.Logf("Calling unpacking(%s), result %s\n", testCase.str, result)
		//Assert
		if testCase.expected != result && err == nil {
			t.Errorf("Incorrect result, Exept %s, got %s", testCase.expected, result)
		}
	}
}
