package tx_temp_cont1_600_032

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/hkraftno/lora-payload-parsers/binary_utils"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(hexString string) (jsonData []byte, err error) {
	var instance Tx_Temp_Cont1_600_032
	err = instance.Load(hexString)
	if err != nil {
		return nil, err
	}

	return json.Marshal(instance)
}

type Tx_Temp_Cont1_600_032 struct {
	ID                uint32      `json:"id"`
	SensorType        uint8       `json:"sensor_type"`
	SequentialCounter uint8       `json:"sequential_counter"`
	FirmWare          uint8       `json:"firm_ware"`
	Settings          uint8       `json:"settings"`
	Temperature       float32     `json:"temperature"`
	Temperature2      float32     `json:"temperature2"`
	AlarmStatus       AlarmStatus `json:"alarm_status"`
	Status            Status      `json:"status"`
}

func (t *Tx_Temp_Cont1_600_032) Load(hexString string) error {
	var sensorType uint8 = 15
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("could not parse hex string as hex: %w", err)
	} else if len(hexBytes) != 14 {
		return fmt.Errorf("hex wrong length to parse, needs to be 22 bytes")
	} else if hexBytes[3] != sensorType {
		return fmt.Errorf("this parser is only for sensor type %d (was %d)", sensorType, hexBytes[3])
	}

	t.ID = binary_utils.BytesToUint32(hexBytes[0:3]) // ID is 3 bytes (6 hex chars)

	t.SensorType = hexBytes[3]
	t.SequentialCounter = hexBytes[4]
	t.FirmWare = hexBytes[5] & 0x03 // first 3 bits are firmware
	t.Settings = hexBytes[5] >> 3   // the rest of the byte is settings
	temp1Int := int16(hexBytes[6])<<8|int16(hexBytes[7])
	temp2Int := int16(hexBytes[8])<<8|int16(hexBytes[9])
	t.Temperature = float32(temp1Int) / 10.0
	t.Temperature2 = float32(temp2Int) / 10.0
	t.AlarmStatus = toAlarmStatus(hexBytes[10:12])
	t.Status = toStatus(hexBytes[12:14])
	return nil
}

func toAlarmStatus(hex []byte) AlarmStatus {
	return AlarmStatus{ // LSB 0 bit numbering:
		TempHiOK: hex[1]&0x01 > 0, // bit 0
		TempLoOK: hex[1]&0x02 > 0, // bit 1
	}
}

type AlarmStatus struct {
	TempHiOK bool `json:"temp_hi_ok"`
	TempLoOK bool `json:"temp_lo_ok"`
}

func toStatus(hex []byte) Status {
	return Status{ // LSB 0 bit numbering:
		MessageTypeIsAlarm: hex[1]&0x01 > 0,                         // bit 0
		BatteryPercentage:  toBatteryPercentage(hex[1] & 0x0C >> 2), // bit 2 and 3
	}
}

type Status struct {
	MessageTypeIsAlarm bool  `json:"msg_type_is_alarm"`
	BatteryPercentage  uint8 `json:"battery_percentage"`
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
