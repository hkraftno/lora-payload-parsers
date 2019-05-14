package pul_lab_xxns

import (
	"math"
	"testing"
)

func TestID(t *testing.T) {
	var expected uint8 = 2
	var data PulLabxxnsStruct
	data.parse([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00})
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
	var data PulLabxxnsStruct
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
	var data PulLabxxnsStruct
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

func TestWireCutEnabled(t *testing.T) {
	var expected = true
	var data PulLabxxnsStruct
	data.parse([]byte{0x07, 0x80, 0x00, 0x00, 0x00, 0x00})
	actual := data.WireCutStatus
	if expected != *actual {
		t.Errorf(
			"Expected WireCutStatus to be %t but was %t",
			expected,
			*actual,
		)
	}
}

func TestWireCutDisabled(t *testing.T) {
	var data PulLabxxnsStruct
	data.parse([]byte{0x03, 0xfd, 0x00, 0x00, 0x00, 0x00})
	actual := data.WireCutStatus
	if nil != actual {
		t.Errorf(
			"Expected WireCutStatus to be nil but was %v",
			actual,
		)
	}
}

func TestInternalData(t *testing.T) {
	expected := "01020304050607"
	var data PulLabxxnsStruct
	data.parse([]byte{0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x00, 0x00, 0x00, 0x00})
	actual := data.InternalData
	if actual != expected {
		t.Errorf(
			"Expected InternalData to be %s but was %s",
			expected,
			actual,
		)
	}
}

func TestCounter(t *testing.T) {
	var expected uint32 = 2147483647
	var data PulLabxxnsStruct
	data.parse([]byte{0x00, 0x00, 0x7f, 0xff, 0xff, 0xff})
	actual := data.Counter
	if expected != actual {
		t.Errorf(
			"Expected Counter to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestCounterVariableInternalData(t *testing.T) {
	var expected uint32 = 16843009
	var data PulLabxxnsStruct
	data.parse([]byte{0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01})
	actual := data.Counter
	if expected != actual {
		t.Errorf(
			"Expected Counter to be %d but was %d",
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
	parsedPayload, err := Parse([]byte{0x02, 0xfc, 0x8e, 0x01, 0x9c, 0x10, 0x00, 0x1b, 0x1b, 0x00})
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	var expected = `{"id":2,"wire_cut_status":null,"battery_level":99,"internal_data":"8e019c10","counter":1776384}`
	if expected != string(parsedPayload) {
		t.Errorf(
			"Expected JSON to be\n%s\nbut was\n%s",
			expected,
			parsedPayload,
		)
	}
}
