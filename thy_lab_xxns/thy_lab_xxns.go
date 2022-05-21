package thy_lab_xxns

import (
	"encoding/json"
	"fmt"
	"math"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(payload []byte) (jsonData []byte, err error) {
	if payload[0] != 0x03 {
		return nil, fmt.Errorf("this parser only supports the Uplink message DATALOG - FPort 3 (HEX starts with 03)")
	}
	var data ThyLabxxnsStruct
	data.Load(payload)

	return json.Marshal(data)
}

type ThyLabxxnsStruct struct {
	ID           uint8   `json:"id"`
	BatteryLevel uint8   `json:"battery_level"`
	InternalData string  `json:"internal_data"`
	Temperature  float32 `json:"temperature"`
	Humidity     uint8   `json:"humidity"`
}

func (t *ThyLabxxnsStruct) Load(payload []byte) {
	length := len(payload)
	t.ID = uint8(payload[0])
	// battery level expressed in 1/254 %
	t.BatteryLevel = uint8(math.Round(float64(payload[1]) / 254.0 * 100))
	for _, b := range payload[2 : length-3] {
		t.InternalData += fmt.Sprintf("%02x", b)
	}
	msb := int16(payload[length-3])
	lsb := int16(payload[length-2])
	// temperature expressed in 1/16 °C as a 2 bytes signed int
	t.Temperature = float32((msb<<8)|lsb) / 16
	t.Humidity = uint8(payload[length-1])
}
