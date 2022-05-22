package ope_lab_xxns

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(hexString string) (jsonData []byte, err error) {
	var data OpeLabxxnsStruct
	err = data.Load(hexString)
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

type OpeLabxxnsStruct struct {
	ID           uint8  `json:"id"`
	BatteryLevel uint8  `json:"battery_level"`
	OpenState    bool   `json:"open_state"`
	InternalData string `json:"internal_data"`
}

func (t *OpeLabxxnsStruct) Load(hexString string) error {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("could not parse hex string as hex: %w", err)
	} else if hexBytes[0] != 0x05 {
		return fmt.Errorf("this parser only supports the Uplink message DATALOG - FPort 3 (HEX starts with 05)")
	} else if len(hexBytes) < 3 {
		return fmt.Errorf("hex too short to parse, needs at least 3 bytes")
	}
	t.ID = uint8(hexBytes[0])
	// battery level expressed in 1/254 %
	t.BatteryLevel = uint8(math.Round(float64(hexBytes[1]) / 254.0 * 100))
	t.OpenState = hexBytes[2]&0x80 > 0
	for _, b := range hexBytes[3:] {
		t.InternalData += fmt.Sprintf("%02x", b)
	}
	return nil
}
