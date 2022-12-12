package tx_temp_cont1_600_032

import (
	"testing"
)

func TestErr(t *testing.T) {
	var data Tx_Temp_Cont1_600_032
	err := data.Load("05")
	if err == nil {
		t.Errorf("Expected error to occur because hex string was too short")
	}
}

func TestID(t *testing.T) {
	var expected uint32 = 220
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0000000000000")
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
	expected := 15 // Tx Temp Cont1
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0000000000000")
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
	expected := 10
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0000000000000")
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
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0000000000000")
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
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0900D0000000000000")
	actual := data.Settings
	if actual != expected {
		t.Errorf(
			"Expected Settins to be %d but was %d",
			expected,
			actual,
		)
	}
}

func TestTemperature(t *testing.T) {
	var expected float32 = 20.8
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0000000000000")
	actual := data.Temperature
	if expected != actual {
		t.Errorf(
			"Expected Temperature to be %.1f but was %.1f",
			expected,
			actual,
		)
	}
}

func TestNegativeTemperature(t *testing.T) {
	var expected float32 = -0.1
	var data Tx_Temp_Cont1_600_032
	data.Load("0050520f6c09ffff000000000000")
	actual := data.Temperature
	if expected != actual {
		t.Errorf(
			"Expected Temperature to be %.1f but was %.1f",
			expected,
			actual,
		)
	}
}

func TestTemperature2(t *testing.T) {
	var expected float32 = 14.4
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0009000000000")
	actual := data.Temperature2
	if expected != actual {
		t.Errorf(
			"Expected PulseCh2 to be %.1f but was %.1f",
			expected,
			actual,
		)
	}
}

func TestAlarmStatus(t *testing.T) {
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0000000020000")
	tests := []struct {
		variableName string
		variable     any
		expected     any
	}{
		{"AlarmStatus.TempHiOK", data.AlarmStatus.TempHiOK, false},
		{"AlarmStatus.TempLoOK", data.AlarmStatus.TempLoOK, true},
	}
	for _, testData := range tests {
		if testData.variable != testData.expected {
			t.Errorf("Expected %s to be %v (was %v)", testData.variableName, testData.expected, testData.variable)
		}
	}
}

func TestStatus(t *testing.T) {
	var data Tx_Temp_Cont1_600_032
	data.Load("0000DC0F0A0100D0000000000001")
	tests := []struct {
		variableName string
		variable     any
		expected     any
	}{
		{"Status.MessageTypeIsAlarm", data.Status.MessageTypeIsAlarm, true},
	}
	for _, testData := range tests {
		if testData.variable != testData.expected {
			t.Errorf("Expected %s to be %v (was %v)", testData.variableName, testData.expected, testData.variable)
		}
	}
}

func TestStatusBatteryPercentage(t *testing.T) {
	var data Tx_Temp_Cont1_600_032
	tests := []struct {
		hex          string
		variableName string
		variable     func(Tx_Temp_Cont1_600_032) uint8
		expected     uint8
	}{
		{"0000DC0F0A0100D0000000000021", "status.BatteryPercentage", func(data Tx_Temp_Cont1_600_032) uint8 { return data.Status.BatteryPercentage }, 100},
		{"0000DC0F0A0100D0000000000025", "status.BatteryPercentage", func(data Tx_Temp_Cont1_600_032) uint8 { return data.Status.BatteryPercentage }, 75},
		{"0000DC0F0A0100D0000000000029", "status.BatteryPercentage", func(data Tx_Temp_Cont1_600_032) uint8 { return data.Status.BatteryPercentage }, 50},
		{"0000DC0F0A0100D000000000002D", "status.BatteryPercentage", func(data Tx_Temp_Cont1_600_032) uint8 { return data.Status.BatteryPercentage }, 25},
	}
	for _, testData := range tests {
		data.Load(testData.hex)
		if testData.variable(data) != testData.expected {
			t.Errorf("Expected %s to be %v (was %v)", testData.variableName, testData.expected, testData.variable(data))
		}
	}
}
