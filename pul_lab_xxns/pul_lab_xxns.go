package pul_lab_xxns

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/hkraftno/hkraft-iot/lora-payload-parsers/binary_utils"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(payload []byte) (jsonData []byte, err error) {
	if !(payload[0] == 02 || payload[0] == 07) {
		return nil, fmt.Errorf("this parser only supports the Uplink message DATALOG - FPort 3 (HEX starts with 02 or 07)")
	}
	var data PulLabxxnsStruct
	data.Load(payload)

	return json.Marshal(data)
}

type PulLabxxnsStruct struct {
	ID            uint8  `json:"id"`
	WireCutStatus *bool  `json:"wire_cut_status"`
	BatteryLevel  uint8  `json:"battery_level"`
	InternalData  string `json:"internal_data"`
	Counter       uint32 `json:"counter"`
}

func (t *PulLabxxnsStruct) Load(payload []byte) {
	length := len(payload)
	t.ID = uint8(payload[0])
	batteryLocation := 1

	if t.ID == 7 {
		wirecut := payload[1]&0x80 > 0
		t.WireCutStatus = &wirecut
		batteryLocation = 2
	}
	// battery level expressed in 1/254 %
	t.BatteryLevel = uint8(math.Round(float64(payload[batteryLocation]) / 254.0 * 100))
	for _, b := range payload[2 : length-4] {
		t.InternalData += fmt.Sprintf("%02x", b)
	}
	// detection number as a 4 bytes unsigned int
	t.Counter = binary_utils.BytesToUint32(payload[length-4 : length])
}
