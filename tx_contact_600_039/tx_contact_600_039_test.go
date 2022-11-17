package tx_contact_600_039

import (
	"testing"
)

func TestErr(t *testing.T) {
	var data Tx_Contact_600_039
	err := data.Load("05")
	if err == nil {
		t.Errorf("Expected error to occur because hex string was too short")
	}
}

func TestID(t *testing.T) {
	var expected uint32 = 13619
	var data Tx_Contact_600_039
	data.Load("0035330bde0900000000000000000000000000000000")
	actual := data.ID
	if expected != actual {
		t.Errorf(
			"Expected ID to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestSensorType(t *testing.T) {
	expected := 11 // Tx Contact
	var data Tx_Contact_600_039
	data.Load("0035330b180900000000000000000000000000000000")
	actual := data.SensorType
	if uint8(expected) != actual {
		t.Errorf(
			"Expected sensor type to be %d but was %d",
			uint8(expected),
			actual,
		)
	}
}

func TestSequentialCounter(t *testing.T) {
	expected := 24
	var data Tx_Contact_600_039
	data.Load("0035330b180900000000000000000000000000000000")
	actual := data.SequentialCounter
	if uint8(expected) != actual {
		t.Errorf(
			"Expected Sequential Counter to be %d but was %d",
			uint8(expected),
			actual,
		)
	}
}

func TestFirmWare(t *testing.T) {
	var expected uint8 = 1
	var data Tx_Contact_600_039
	data.Load("0035330b180900000000000000000000000000000000")
	actual := data.FirmWare
	if actual != expected {
		t.Errorf(
			"Expected FirmWare to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestSettings(t *testing.T) {
	var expected uint8 = 1
	var data Tx_Contact_600_039
	data.Load("0035330b180900000000000000000000000000000000")
	actual := data.Settings
	if actual != expected {
		t.Errorf(
			"Expected Settins to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestPulseCh1(t *testing.T) {
	var expected uint32 = 1
	var data Tx_Contact_600_039
	data.Load("0035330b180900000001000000010000000000000000")
	actual := data.PulseCh1
	if expected != actual {
		t.Errorf(
			"Expected PulseCh1 to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestPulseCh2(t *testing.T) {
	var expected uint32 = 1
	var data Tx_Contact_600_039
	data.Load("0035330b180900000001000000010000000000000000")
	actual := data.PulseCh2
	if expected != actual {
		t.Errorf(
			"Expected PulseCh2 to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestPulseOC(t *testing.T) {
	var expected uint32 = 1
	var data Tx_Contact_600_039
	data.Load("0035330b180900000001000000010000000100000000")
	actual := data.PulseCh2
	if expected != actual {
		t.Errorf(
			"Expected PulseOC to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestAlarmStatus(t *testing.T) {
	var data Tx_Contact_600_039
	data.Load("0035330b710900000005000000000000000000010021")
	tests := []struct {
		variableName string
		variable     any
		expected     any
	}{
		{"AlarmStatus.PulseCh1Change", data.AlarmStatus.PulseCh1Change, true},
		{"AlarmStatus.PulseCh2Change", data.AlarmStatus.PulseCh2Change, false},
		{"AlarmStatus.PulseOCChange", data.AlarmStatus.PulseOCChange, false},
	}
	for _, testData := range tests {
		if testData.variable != testData.expected {
			t.Errorf("Expected %s to be %v (was %v)", testData.variableName, testData.expected, testData.variable)
		}
	}
}

func TestStatus(t *testing.T) {
	var data Tx_Contact_600_039
	data.Load("0035330b710900000005000000000000000000010021")
	tests := []struct {
		variableName string
		variable     any
		expected     any
	}{
		{"Status.MessageTypeIsAlarm", data.Status.MessageTypeIsAlarm, true},
		{"Status.PulseCh1IsClosed", data.Status.PulseCh1IsClosed, true},
		{"Status.PulseCh2IsClosed", data.Status.PulseCh2IsClosed, false},
		{"Status.PulseOCIsClosed", data.Status.PulseOCIsClosed, false},
	}
	for _, testData := range tests {
		if testData.variable != testData.expected {
			t.Errorf("Expected %s to be %v (was %v)", testData.variableName, testData.expected, testData.variable)
		}
	}
}

func TestStatusBatteryPercentage(t *testing.T) {
	var data Tx_Contact_600_039
	tests := []struct {
		hex          string
		variableName string
		variable     func(Tx_Contact_600_039) uint8
		expected     uint8
	}{
		{"0035330b710900000005000000000000000000010021", "status.BatteryPercentage", func(data Tx_Contact_600_039) uint8 { return data.Status.BatteryPercentage }, 100},
		{"0035330b710900000005000000000000000000010025", "status.BatteryPercentage", func(data Tx_Contact_600_039) uint8 { return data.Status.BatteryPercentage }, 75},
		{"0035330b710900000005000000000000000000010029", "status.BatteryPercentage", func(data Tx_Contact_600_039) uint8 { return data.Status.BatteryPercentage }, 50},
		{"0035330b71090000000500000000000000000001002D", "status.BatteryPercentage", func(data Tx_Contact_600_039) uint8 { return data.Status.BatteryPercentage }, 25},
	}
	for _, testData := range tests {
		data.Load(testData.hex)
		if testData.variable(data) != testData.expected {
			t.Errorf("Expected %s to be %v (was %v)", testData.variableName, testData.expected, testData.variable(data))
		}
	}
}
