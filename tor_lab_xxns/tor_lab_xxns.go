package tor_lab_xxns

import (
	"encoding/json"
	"fmt"
	"math"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(payload []byte) (jsonData []byte, err error) {
	if payload[0] != 0x05 {
		return nil, fmt.Errorf("This parser only supports the Uplink message DATALOG - FPort 3 (HEX starts with 05)")
	}
	var data TorLabxxnsStruct
	data.parse(payload)

	return json.Marshal(data)
}

type TorLabxxnsStruct struct {
	ID           uint8  `json:"id"`
	BatteryLevel uint8  `json:"battery_level"`
	OpenState    bool   `json:"open_state"`
	InternalData string `json:"internal_data"`
}

func (t *TorLabxxnsStruct) parse(payload []byte) {
	t.ID = uint8(payload[0])
	// battery level expressed in 1/254 %
	t.BatteryLevel = uint8(math.Round(float64(payload[1]) / 254.0 * 100))
	t.OpenState = payload[2]&0x80 > 0
	for _, b := range payload[3:] {
		t.InternalData += fmt.Sprintf("%02x", b)
	}
}
