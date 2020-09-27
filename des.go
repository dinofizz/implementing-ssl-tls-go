package tlsimpl

// Does not return a 1 for a 1 bit, it just returns non-zero
func getBit(array []byte, bit uint8) uint8 {
	location := bit / 8;
	return uint8(array[location] & (0x80 >> (bit % 8)))
}

func setBit(array *[]byte, bit uint8) {
	location := bit / 8;
	(*array)[location] = (*array)[location] | (0x80 >> (bit % 8))
}

func clearBit(array *[]byte, bit uint8) {
	location := bit / 8;
	(*array)[location] = (*array)[location] & ^(0x80 >> (bit % 8))
}

func xor(target *[]byte, src *[]byte) {
	for i, a := range(*target) {
		(*target)[i] = a ^ (*src)[i]
	}
}
