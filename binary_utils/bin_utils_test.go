package binary_utils

import "testing"

func TestThatTooLongArrayFails(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected the code to panic, it did not")
		}
	}()
	BytesToUint32([]byte{0x00, 0x00, 0x00, 0x00, 0x00})
}

func TestThatNumbersAreParsedCorrectly(t *testing.T) {
	var expected uint32 = 6554659
	number := BytesToUint32([]byte{0x64, 0x04, 0x23})
	if number != expected {
		t.Errorf("Expected number to be %d but was %d", number, expected)
	}
}
