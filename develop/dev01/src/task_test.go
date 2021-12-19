package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		host string
		err  error
	}{
		{
			host: "0.beevik-ntp.pool.ntp.org",
			err:  nil,
		},
		{
			host: "0.wrong.host.ntp.org",
			err:  nil,
		},
	}
	//Act
	for _, testCase := range testTable {
		result, err := getTime(testCase.host)
		t.Logf("Calling getTime(%s), result %s", testCase.host, result)
		//Assert
		if err != nil {
			t.Errorf("end with err %v", err)
		}

	}
}
