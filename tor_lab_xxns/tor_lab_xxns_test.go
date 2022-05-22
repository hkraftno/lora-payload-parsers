package tor_lab_xxns

import (
	"math"
	"testing"
)

func TestErr(t *testing.T) {
	var data TorLabxxnsStruct
	err := data.Load("05")
	if err == nil {
		t.Errorf("Expected error to occur because hex string was too short")
	}
}

func TestID(t *testing.T) {
	var expected uint8 = 5
	var data TorLabxxnsStruct
	data.Load("050000000000")
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
	var data TorLabxxnsStruct
	data.Load("05fc00000000")
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
	var data TorLabxxnsStruct
	data.Load("05fd00000000")
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
	var data TorLabxxnsStruct
	data.Load("05000001020304050607")
	actual := data.InternalData
	if actual != expected {
		t.Errorf(
			"Expected InternalData to be %s but was %s",
			expected,
			actual,
		)
	}
}

func TestOpenStateTrue(t *testing.T) {
	var expected bool = true
	var data TorLabxxnsStruct
	data.Load("050083")
	actual := data.OpenState
	if expected != actual {
		t.Errorf(
			"Expected OpenState to be %t but was %t",
			expected,
			actual,
		)
	}
}

func TestOpenStateFalse(t *testing.T) {
	var expected bool = false
	var data TorLabxxnsStruct
	data.Load("050003")
	actual := data.OpenState
	if expected != actual {
		t.Errorf(
			"Expected OpenState to be %t but was %t",
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
	parsedPayload, err := Parse("05fc808e019c10")
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	var expected = `{"id":5,"battery_level":99,"open_state":true,"internal_data":"8e019c10"}`
	if expected != string(parsedPayload) {
		t.Errorf(
			"Expected JSON to be\n%s\nbut was\n%s",
			expected,
			parsedPayload,
		)
	}
}
