package thy_lab_xxns

import (
	"math"
	"testing"
)

func TestID(t *testing.T) {
	var expected uint8 = 3
	var data ThyLabxxnsStruct
	data.Load([]byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00})
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
	var data ThyLabxxnsStruct
	data.Load([]byte{0x00, 0xfc, 0x00, 0x00, 0x00, 0x00})
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
	var data ThyLabxxnsStruct
	data.Load([]byte{0x00, 0xfd, 0x00, 0x00, 0x00, 0x00})
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
	var data ThyLabxxnsStruct
	data.Load([]byte{0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x00, 0x00, 0x00})
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
	var data ThyLabxxnsStruct
	data.Load([]byte{0x00, 0x00, 0x7f, 0xff, 0x00})
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
	var data ThyLabxxnsStruct
	data.Load([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xff, 0x00})
	actual := data.Temperature
	if expected != actual {
		t.Errorf(
			"Expected Temperature to be %f but was %f",
			expected,
			actual,
		)
	}
}

func TestHumidity(t *testing.T) {
	var expected uint8 = 95
	var data ThyLabxxnsStruct
	data.Load([]byte{0x00, 0x00, 0x00, 0x00, 0x5f})
	actual := data.Humidity
	if expected != actual {
		t.Errorf(
			"Expected Humidity to be %d but was %d",
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
	parsedPayload, err := Parse([]byte{0x03, 0xfc, 0x8e, 0x01, 0x9c, 0x10, 0x00, 0x1b, 0x63})
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	var expected = `{"id":3,"battery_level":99,"internal_data":"8e019c10","temperature":1.6875,"humidity":99}`
	if expected != string(parsedPayload) {
		t.Errorf(
			"Expected JSON to be\n%s\nbut was\n%s",
			expected,
			parsedPayload,
		)
	}
}
