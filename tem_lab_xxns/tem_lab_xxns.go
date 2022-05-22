package tem_lab_xxns

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(hexString string) (jsonData []byte, err error) {
	var data TemLabxxnsStruct
	err = data.Load(hexString)
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

type TemLabxxnsStruct struct {
	ID           uint8   `json:"id"`
	BatteryLevel uint8   `json:"battery_level"`
	InternalData string  `json:"internal_data"`
	Temperature  float32 `json:"temperature"`
}

func (t *TemLabxxnsStruct) Load(hexString string) error {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("could not parse hex string as hex: %w", err)
	} else if hexBytes[0] != 0x01 {
		return fmt.Errorf("this parser only supports the Uplink message DATALOG - FPort 3 (HEX starts with 01)")
	} else if len(hexBytes) < 4 {
		return fmt.Errorf("hex too short to parse, needs at least 4 bytes")
	}
	length := len(hexBytes)
	t.ID = uint8(hexBytes[0])
	// battery level expressed in 1/254 %
	t.BatteryLevel = uint8(math.Round(float64(hexBytes[1]) / 254.0 * 100))
	for _, b := range hexBytes[2 : length-2] {
		t.InternalData += fmt.Sprintf("%02x", b)
	}
	msb := int16(hexBytes[length-2])
	lsb := int16(hexBytes[length-1])
	// temperature expressed in 1/16 °C as a 2 bytes signed int
	t.Temperature = float32((msb<<8)|lsb) / 16
	return nil
}
