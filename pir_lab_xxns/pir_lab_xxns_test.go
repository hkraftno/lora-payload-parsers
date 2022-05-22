package pir_lab_xxns

import (
	"math"
	"testing"
)

func TestErr(t *testing.T) {
	var data PirLabxxnsStruct
	err := data.Load("05")
	if err == nil {
		t.Errorf("Expected error to occur because hex string was too short")
	}
}


func TestID(t *testing.T) {
	var expected uint8 = 2
	var data PirLabxxnsStruct
	data.Load("020000000000")
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
	var data PirLabxxnsStruct
	data.Load("02fc00000000")
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
	var data PirLabxxnsStruct
	data.Load("02fd00000000")
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
	var data PirLabxxnsStruct
	data.Load("02000102030405060700000000")
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
	var data PirLabxxnsStruct
	data.Load("02007fffffff")
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
	var data PirLabxxnsStruct
	data.Load("0200000001010101")
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
	_, err := Parse("99")
	if err == nil {
		t.Errorf("Expected error when parsing invalid message type, got nil")
	}
}

func TestParserExampleHex1(t *testing.T) {

	parsedPayload, err := Parse("02fc8e019c10001b1b00")
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	var expected = `{"id":2,"battery_level":99,"internal_data":"8e019c10","counter":1776384}`
	if expected != string(parsedPayload) {
		t.Errorf(
			"Expected JSON to be\n%s\nbut was\n%s",
			expected,
			parsedPayload,
		)
	}
}
