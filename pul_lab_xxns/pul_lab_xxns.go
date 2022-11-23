package pul_lab_xxns

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"

	"github.com/hkraftno/lora-payload-parsers/binary_utils"
)

// Parse is the cloud function for converting the payload hex to json
func Parse(hexString string) (jsonData []byte, err error) {
	var data PulLabxxnsStruct
	err = data.Load(hexString)
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

type PulLabxxnsStruct struct {
	ID            uint8  `json:"id"`
	WireCutStatus *bool  `json:"wire_cut_status"`
	BatteryLevel  uint8  `json:"battery_level"`
	InternalData  string `json:"internal_data"`
	Counter       uint32 `json:"counter"`
}

func (t *PulLabxxnsStruct) Load(hexString string) error {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("could not parse hex string as hex: %w", err)
	} else if !(hexBytes[0] == 0x02 || hexBytes[0] == 0x07) {
		return fmt.Errorf("this parser only supports the Uplink message DATALOG - FPort 3 (HEX starts with 02 or 07)")
	} else if hexBytes[0] == 0x02 && len(hexBytes) < 6 {
		return fmt.Errorf("hex too short to parse, needs at least 6 bytes")
	} else if hexBytes[0] == 0x07 && len(hexBytes) < 7 {
		return fmt.Errorf("hex too short to parse, needs at least 7 bytes")
	}
	length := len(hexBytes)
	t.ID = uint8(hexBytes[0])
	batteryLocation := 1

	if t.ID == 7 {
		wirecut := hexBytes[1]&0x80 > 0
		t.WireCutStatus = &wirecut
		batteryLocation = 2
	}
	// battery level expressed in 1/254 %
	t.BatteryLevel = uint8(math.Round(float64(hexBytes[batteryLocation]) / 254.0 * 100))
	for _, b := range hexBytes[2 : length-4] {
		t.InternalData += fmt.Sprintf("%02x", b)
	}
	// detection number as a 4 bytes unsigned int
	t.Counter = binary_utils.BytesToUint32(hexBytes[length-4 : length])
	return nil
}
