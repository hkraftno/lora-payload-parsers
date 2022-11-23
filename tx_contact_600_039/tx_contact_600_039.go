package tx_contact_600_039

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/hkraftno/lora-payload-parsers/binary_utils"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(hexString string) (jsonData []byte, err error) {
	var instance Tx_Contact_600_039
	err = instance.Load(hexString)
	if err != nil {
		return nil, err
	}

	return json.Marshal(instance)
}

type Tx_Contact_600_039 struct {
	ID                uint32      `json:"id"`
	SensorType        uint8       `json:"sensor_type"`
	SequentialCounter uint8       `json:"sequential_counter"`
	FirmWare          uint8       `json:"firm_ware"`
	Settings          uint8       `json:"settings"`
	PulseCh1          uint32      `json:"pulse_ch1"`
	PulseCh2          uint32      `json:"pulse_ch2"`
	PulseOC           uint32      `json:"pulse_oc"`
	AlarmStatus       AlarmStatus `json:"alarm_status"`
	Status            Status      `json:"status"`
}

func (t *Tx_Contact_600_039) Load(hexString string) error {
	var sensorType uint8 = 11
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("could not parse hex string as hex: %w", err)
	} else if len(hexBytes) != 22 {
		return fmt.Errorf("hex too short to parse, needs at least 22 bytes")
	} else if hexBytes[3] != sensorType {
		return fmt.Errorf("this parser is only for sensor type %d (was %d)", sensorType, hexBytes[3])
	}
	t.ID = binary_utils.BytesToUint32(hexBytes[0:3]) // ID is 3 bytes (6 hex chars)

	t.SensorType = hexBytes[3]
	t.SequentialCounter = hexBytes[4]
	t.FirmWare = hexBytes[5] & 0x03 // first 3 bits are firmware
	t.Settings = hexBytes[5] >> 3   // the rest of the byte is settings
	t.PulseCh1 = binary_utils.BytesToUint32(hexBytes[6:10])
	t.PulseCh2 = binary_utils.BytesToUint32(hexBytes[10:14])
	t.PulseOC = binary_utils.BytesToUint32(hexBytes[14:18])
	t.AlarmStatus = toAlarmStatus(hexBytes[18:20])
	t.Status = toStatus(hexBytes[20:22])
	return nil
}

func toAlarmStatus(hex []byte) AlarmStatus {
	return AlarmStatus{ // LSB 0 bit numbering:
		PulseCh1Change: hex[1]&0x01 > 0, // bit 0
		PulseCh2Change: hex[1]&0x02 > 0, // bit 1
		PulseOCChange:  hex[1]&0x04 > 0, // bit 2
	}
}

type AlarmStatus struct {
	PulseCh1Change bool `json:"pulse_ch1_change"`
	PulseCh2Change bool `json:"pulse_ch2_change"`
	PulseOCChange  bool `json:"pulse_oc_change"`
}

func toStatus(hex []byte) Status {
	return Status{ // LSB 0 bit numbering:
		MessageTypeIsAlarm: hex[1]&0x01 > 0,                         // bit 0
		BatteryPercentage:  toBatteryPercentage(hex[1] & 0x0C >> 2), // bit 2 and 3
		PulseCh1IsClosed:   hex[1]&0x20 > 0,                         // bit 5
		PulseCh2IsClosed:   hex[1]&0x40 > 0,                         // bit 6
		PulseOCIsClosed:    hex[1]&0x80 > 0,                         // bit 7
	}
}

type Status struct {
	MessageTypeIsAlarm bool  `json:"msg_type_is_alarm"`
	BatteryPercentage  uint8 `json:"battery_percentage"`
	PulseCh1IsClosed   bool  `json:"pulse_ch1_is_closed"`
	PulseCh2IsClosed   bool  `json:"pulse_ch2_is_closed"`
	PulseOCIsClosed    bool  `json:"pulse_oc_is_closed"`
}

func toBatteryPercentage(hex uint8) uint8 {
	switch hex {
	case 0x00: // 00
		return 100
	case 0x01: // 01
		return 75
	case 0x02: // 10
		return 50
	default: // 0x03 ==> 11
		return 25
	}
}
