package common

// Does not return a 1 for a 1 bit, it just returns non-zero
func GetBit(array []byte, bit int) uint8 {
	location := bit / 8
	return uint8(array[location] & (0x80 >> (bit % 8)))
}

func SetBit(array []byte, bit int) {
	location := bit / 8
	array[location] = array[location] | (0x80 >> (bit % 8))
}

func ClearBit(array []byte, bit int) {
	location := bit / 8
	array[location] = array[location] & ^(0x80 >> (bit % 8))
}

func Xor(target []byte, src []byte, length int) {
	for i := 0; i < length; i++ {
		target[i] = target[i] ^ src[i]
	}
}

func InvertByte(b byte) byte {
	return 0xFF ^ b
}
