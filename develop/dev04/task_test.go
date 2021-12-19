package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := struct {
		str      []string
		expected map[string][]string
	}{
		str:      []string{"пятак", "листок", "пятка", "тяпка", "слиток", "столик", "авб", "Go", "OG", "Ы", "РОМ", "Один", "мор", "бав", "абв", "ваб", "актяп"},
		expected: map[string][]string{},
	}
	testTable.expected["пятак"] = []string{"актяп", "пятак", "пятка", "тяпка"}
	testTable.expected["go"] = []string{"go", "og"}
	testTable.expected["авб"] = []string{"абв", "авб", "бав", "ваб"}
	testTable.expected["листок"] = []string{"листок", "слиток", "столик"}
	testTable.expected["ром"] = []string{"мор", "ром"}
	//Act
	result := findAnagrams(testTable.str)
	t.Logf("Calling manSort(%s), result %s\n", testTable.str, result)
	//Assert
	for k, v := range testTable.expected {
		res := result[k]
		for key, value := range v {
			if value != res[key] {
				t.Errorf("Incorrect result, Exept %s, got %s", testTable.expected, result)
				break
			}
		}

	}

}
