package tem_lab_xxns

import (
	"math"
	"testing"
)

func TestErr(t *testing.T) {
	var data TemLabxxnsStruct
	err := data.Load("05")
	if err == nil {
		t.Errorf("Expected error to occur because hex string was too short")
	}
}

func TestID(t *testing.T) {
	var expected uint8 = 1
	var data TemLabxxnsStruct
	data.Load("010000000000")
	actual := data.ID
	if expected != actual {
		t.Errorf(
			"Expected ID to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestBattery(t *testing.T) {
	expected := math.Round(252.0 / 254.0 * 100) // 99%
	var data TemLabxxnsStruct
	data.Load("01fc00000000")
	actual := data.BatteryLevel
	if uint8(expected) != actual {
		t.Errorf(
			"Expected BatteryLevel to be %d but was %d",
			uint8(expected),
			actual,
		)
	}
}

func TestBatteryRounding(t *testing.T) {
	expected := math.Round(253.0 / 254.0 * 100) // 100%
	var data TemLabxxnsStruct
	data.Load("01fd00000000")
	actual := data.BatteryLevel
	if uint8(expected) != actual {
		t.Errorf(
			"Expected BatteryLevel to be %d but was %d",
			uint8(expected),
			actual,
		)
	}
}

func TestInternalData(t *testing.T) {
	expected := "01020304050607"
	var data TemLabxxnsStruct
	data.Load("0100010203040506070000")
	actual := data.InternalData
	if actual != expected {
		t.Errorf(
			"Expected InternalData to be %s but was %s",
			expected,
			actual,
		)
	}
}

func TestTemperature(t *testing.T) {
	var expected float32 = 32767 / 16.0
	var data TemLabxxnsStruct
	data.Load("01007fff")
	actual := data.Temperature
	if expected != actual {
		t.Errorf(
			"Expected Temperature to be %f but was %f",
			expected,
			actual,
		)
	}
}

func TestTemperatureVariableInternalData(t *testing.T) {
	var expected float32 = -3841 / 16.0
	var data TemLabxxnsStruct
	data.Load("010000000000f0ff")
	actual := data.Temperature
	if expected != actual {
		t.Errorf(
			"Expected Temperature to be %f but was %f",
			expected,
			actual,
		)
	}
}

func TestInvalidMessageFormat(t *testing.T) {
	_, err := Parse("99")
	if err == nil {
		t.Errorf("Expected error when parsing invalid message type, got nil")
	}
}

func TestParserExampleHex1(t *testing.T) {
	parsedPayload, err := Parse("01fc8e019c10001b")
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	var expected = `{"id":1,"battery_level":99,"internal_data":"8e019c10","temperature":1.6875}`
	if expected != string(parsedPayload) {
		t.Errorf(
			"Expected JSON to be\n%s\nbut was\n%s",
			expected,
			parsedPayload,
		)
	}
}
