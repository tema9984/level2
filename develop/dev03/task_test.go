package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		str      []string
		expected []string
	}{
		{
			str:      []string{"1 computer 1", "1 computer 1", "4 mouse 4", "5 LAPTOP 5", "1 data 6", "12 RedHat 4", "15 laptop 2", "17 debian 5", "11 laptop 1"},
			expected: []string{"1 computer 1", "1 computer 1", "1 data 6", "11 laptop 1", "12 RedHat 4", "15 laptop 2", "17 debian 5", "4 mouse 4", "5 LAPTOP 5"},
		},
		{
			str:      []string{"1 computer 1", "1 computer 1", "4 mouse 4", "5 LAPTOP 5", "1 data 6", "12 RedHat 4", "15 laptop 2", "17 debian 5", "11 Laptop 1"},
			expected: []string{"1 computer 1", "1 data 6", "17 debian 5", "15 laptop 2", "11 Laptop 1", "5 LAPTOP 5", "4 mouse 4", "12 RedHat 4"},
		},
		{
			str:      []string{"1 computer 0", "1 computer 0", "4 mouse 4", "5 LAPTOP 7", "1 data 6", "12 RedHat 3", "15 laptop 2", "17 debian 5", "11 laptop 1"},
			expected: []string{"5 LAPTOP 7", "1 data 6", "17 debian 5", "4 mouse 4", "12 RedHat 3", "15 laptop 2", "11 laptop 1", "1 computer 0"},
		},
	}
	//Default
	result := manSort(testTable[0].str, -1, false, false, false)
	t.Logf("Calling manSort(%s), result %s\n", testTable[0].str, result)
	//Assert
	for k, v := range testTable[0].expected {
		if v != result[k] {
			t.Errorf("Incorrect result, Exept %s, got %s", testTable[0].expected, result)
		}
	}
	//Act
	//column(1) uniq
	result = manSort(testTable[1].str, 1, false, false, true)
	t.Logf("Calling manSort(%s), result %s\n", testTable[1].str, result)
	//Assert
	for k, v := range testTable[1].expected {
		if v != result[k] {
			t.Errorf("Incorrect result, Exept %s, got %s", testTable[1].expected, result)
		}
	}
	//column(2) numer uniq revers
	result = manSort(testTable[2].str, 2, true, true, true)
	t.Logf("Calling manSort(%s), result %s\n", testTable[2].str, result)
	//Assert
	for k, v := range testTable[2].expected {
		if v != result[k] {
			t.Errorf("Incorrect result, Exept %s, got %s", testTable[2].expected, result)
		}
	}

}
