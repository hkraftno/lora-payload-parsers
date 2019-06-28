package ope_lab_xxns

import (
	"math"
	"testing"
)

func TestID(t *testing.T) {
	var expected uint8 = 5
	var data OpeLabxxnsStruct
	data.parse([]byte{0x05, 0x00, 0x00, 0x00, 0x00, 0x00})
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
	var data OpeLabxxnsStruct
	data.parse([]byte{0x00, 0xfc, 0x00, 0x00, 0x00, 0x00})
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
	var data OpeLabxxnsStruct
	data.parse([]byte{0x00, 0xfd, 0x00, 0x00, 0x00, 0x00})
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
	var data OpeLabxxnsStruct
	data.parse([]byte{0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07})
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
	var data OpeLabxxnsStruct
	data.parse([]byte{0x00, 0x00, 0x83})
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
	var data OpeLabxxnsStruct
	data.parse([]byte{0x00, 0x00, 0x03})
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
	_, err := Parse([]byte{0x99})
	if err == nil {
		t.Errorf("Expected error when parsing invalid message type, got nil")
	}
}

func TestParserExampleHex1(t *testing.T) {
	parsedPayload, err := Parse([]byte{0x05, 0xfb, 0x03, 0x15, 0x03, 0x0c, 0x1f})
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	var expected = `{"id":5,"battery_level":99,"open_state":false,"internal_data":"15030c1f"}`
	if expected != string(parsedPayload) {
		t.Errorf(
			"Expected JSON to be\n%s\nbut was\n%s",
			expected,
			parsedPayload,
		)
	}
}
