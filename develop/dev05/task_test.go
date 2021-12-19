package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		str      []string
		expected []string
	}{
		{
			str:      []string{"safasf adgasd", "asagasggasdga", "agah sa ad", "agah sa ad", "hdfhdf gdsg", "adsgasdg a sarha adsg asdg", "asffasds agda sd", "agsdagsa gad asdg asg", "dsfsag asg sadga", "sagdsadg gads a", "hsaad a.hgsdads", "adsgasdg a sarha adsg asdg", "asfagda", "ashdfhsdfh agsdgaga gasdg sadga"},
			expected: []string{"adsgasdg a sarha adsg asdg", "adsgasdg a sarha adsg asdg"},
		},
		{
			str:      []string{"safasf adgasd", "asagasggasdga", "agah sa ad", "agah sa ad", "hdfhdf gdsg", "adsgasdg a sarha adsg asdg", "asffasds agda sd", "agsdagsa gad asdg asg", "dsfsag asg sadga", "sagdsadg gads a", "hsaad a.hgsdads", "adsgasdg a sarha adsg asdg", "asfagda", "ashdfhsdfh agsdgaga gasdg sadga"},
			expected: []string{"2"},
		},
		{
			str:      []string{"safasf adgasd", "asagasggasdga", "agah sa ad", "agah sa ad", "hdfhdf gdsg", "adsgasdg a sarha adsg asdg", "asffasds agda sd", "agsdagsa gad asdg asg", "dsfsag asg sadga", "sagdsadg gads a", "hsaad a.hgsdads", "adsgasdg a sarha adsg asdg", "asfagda", "ashdfhsdfh agsdgaga gasdg sadga"},
			expected: []string{"4:hdfhdf gdsg", "5:adsgasdg a sarha adsg asdg", "6:asffasds agda sd", "10:hsaad a.hgsdads", "11:adsgasdg a sarha adsg asdg", "12:asfagda"},
		},
	}
	//Default
	result := manGreep(testTable[0].str, `ads.*h`, -1, -1, -1, false, false, false, false, false)
	t.Logf("Calling manSort(%s), result %s\n", testTable[0].str, result)
	//Assert
	for k, v := range testTable[0].expected {
		if v != result[k] {
			t.Errorf("Incorrect result, Exept %s, got %s", testTable[0].expected, result)
		}
	}
	//Count
	result = manGreep(testTable[1].str, `ads.*h`, -1, -1, -1, true, false, false, false, false)
	t.Logf("Calling manSort(%s), result %s\n", testTable[1].str, result)
	//Assert
	for k, v := range testTable[1].expected {
		if v != result[k] {
			t.Errorf("Incorrect result, Exept %s, got %s", testTable[1].expected, result)
		}
	}
	//Context(1)+line num
	result = manGreep(testTable[2].str, `ads.*h`, -1, -1, 1, true, false, false, false, true)
	t.Logf("Calling manSort(%s), result %s\n", testTable[2].str, result)
	//Assert
	for k, v := range testTable[2].expected {
		if v != result[k] {
			t.Errorf("Incorrect result, Exept %s, got %s", testTable[2].expected, result)
		}
	}

}
