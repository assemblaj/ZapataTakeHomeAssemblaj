package utils

import "testing"

const testString string = `
Station name not available
Mar 16, 2011 - 01:30 AM EDT / 2011.03.16 0530 UTC
Wind: from the SW (230 degrees) at 5 MPH (4 KT):0
Visibility: greater than 7 mile(s):0
Temperature: 35 F (2 C)
Windchill: 30 F (-1 C):1
Dew Point: 23 F (-5 C)
Relative Humidity: 59%
Pressure (altimeter): 30.39 in. Hg (1029 hPa)
ob: ABBN 160530Z 23004KT 9999 NSC 02/M05 Q1029 R14R/CLRD60 NOSIG RMK G/O QFE696
cycle: 5
`

func TestGetKeyValue(t *testing.T) {
	keyValMap := GetKeyValueFromString(testString)
	if len(keyValMap) == 0 {
		t.Log("Error: Map should not be empty.")
		t.Fail()
	}
	_, hasTemp := keyValMap["Temperature"]

	if !hasTemp {
		t.Log("Error, map should have temperature")
		t.Fail()
	}

	_, hasPressure := keyValMap["Pressure (altimeter)"]
	if !hasPressure {
		t.Log("Error, map should have pressure.", keyValMap)
		t.Fail()
	}
}
