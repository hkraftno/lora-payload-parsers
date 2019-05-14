package binary_utils

func BytesToUint32(byteList []byte) (output uint32) {
	if len(byteList) > 4 {
		panic("Tried to convert byte array greater than 4 to int")
	}
	for _, b := range byteList {
		output = (output << 8) | uint32(b)
	}
	return output
}
