package pir_lab_xxns

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/hkraft/hkraft-iot/lora-payload-parsers/binary_utils"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(payload []byte) (jsonData []byte, err error) {
	if !(payload[0] == 02 || payload[0] == 07) {
		return nil, fmt.Errorf("This parser only supports the Uplink message DATALOG - FPort 3 (HEX starts with 02)")
	}
	var data PirLabxxnsStruct
	data.parse(payload)

	return json.Marshal(data)
}

type PirLabxxnsStruct struct {
	ID           uint8  `json:"id"`
	BatteryLevel uint8  `json:"battery_level"`
	InternalData string `json:"internal_data"`
	Counter      uint32 `json:"counter"`
}

func (t *PirLabxxnsStruct) parse(payload []byte) {
	length := len(payload)
	t.ID = uint8(payload[0])
	// battery level expressed in 1/254 %
	t.BatteryLevel = uint8(math.Round(float64(payload[1]) / 254.0 * 100))
	for _, b := range payload[2 : length-4] {
		t.InternalData += fmt.Sprintf("%02x", b)
	}
	// detection number as a 4 bytes unsigned int
	t.Counter = binary_utils.BytesToUint32(payload[length-4 : length])
}
