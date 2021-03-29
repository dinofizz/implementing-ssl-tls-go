package des

import (
	"log"
)

const DES_BLOCK_SIZE = 8
const DES_KEY_SIZE = 8
const EXPANSION_BLOCK_SIZE = 6
const PC1_KEY_SIZE = 7
const SUBKEY_SIZE = 6

var ipTable = []int{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

var fpTable = []int{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

var pc1Table = []int{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

var pc2Table = []int{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32,
}

var expansionTable = []int{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

var sbox = [][]int{
	{14, 0, 4, 15, 13, 7, 1, 4, 2, 14, 15, 2, 11, 13, 8, 1,
		3, 10, 10, 6, 6, 12, 12, 11, 5, 9, 9, 5, 0, 3, 7, 8,
		4, 15, 1, 12, 14, 8, 8, 2, 13, 4, 6, 9, 2, 1, 11, 7,
		15, 5, 12, 11, 9, 3, 7, 14, 3, 10, 10, 0, 5, 6, 0, 13},
	{15, 3, 1, 13, 8, 4, 14, 7, 6, 15, 11, 2, 3, 8, 4, 14,
		9, 12, 7, 0, 2, 1, 13, 10, 12, 6, 0, 9, 5, 11, 10, 5,
		0, 13, 14, 8, 7, 10, 11, 1, 10, 3, 4, 15, 13, 4, 1, 2,
		5, 11, 8, 6, 12, 7, 6, 12, 9, 0, 3, 5, 2, 14, 15, 9},
	{10, 13, 0, 7, 9, 0, 14, 9, 6, 3, 3, 4, 15, 6, 5, 10,
		1, 2, 13, 8, 12, 5, 7, 14, 11, 12, 4, 11, 2, 15, 8, 1,
		13, 1, 6, 10, 4, 13, 9, 0, 8, 6, 15, 9, 3, 8, 0, 7,
		11, 4, 1, 15, 2, 14, 12, 3, 5, 11, 10, 5, 14, 2, 7, 12},
	{7, 13, 13, 8, 14, 11, 3, 5, 0, 6, 6, 15, 9, 0, 10, 3,
		1, 4, 2, 7, 8, 2, 5, 12, 11, 1, 12, 10, 4, 14, 15, 9,
		10, 3, 6, 15, 9, 0, 0, 6, 12, 10, 11, 1, 7, 13, 13, 8,
		15, 9, 1, 4, 3, 5, 14, 11, 5, 12, 2, 7, 8, 2, 4, 14},
	{2, 14, 12, 11, 4, 2, 1, 12, 7, 4, 10, 7, 11, 13, 6, 1,
		8, 5, 5, 0, 3, 15, 15, 10, 13, 3, 0, 9, 14, 8, 9, 6,
		4, 11, 2, 8, 1, 12, 11, 7, 10, 1, 13, 14, 7, 2, 8, 13,
		15, 6, 9, 15, 12, 0, 5, 9, 6, 10, 3, 4, 0, 5, 14, 3},
	{12, 10, 1, 15, 10, 4, 15, 2, 9, 7, 2, 12, 6, 9, 8, 5,
		0, 6, 13, 1, 3, 13, 4, 14, 14, 0, 7, 11, 5, 3, 11, 8,
		9, 4, 14, 3, 15, 2, 5, 12, 2, 9, 8, 5, 12, 15, 3, 10,
		7, 11, 0, 14, 4, 1, 10, 7, 1, 6, 13, 0, 11, 8, 6, 13},
	{4, 13, 11, 0, 2, 11, 14, 7, 15, 4, 0, 9, 8, 1, 13, 10,
		3, 14, 12, 3, 9, 5, 7, 12, 5, 2, 10, 15, 6, 8, 1, 6,
		1, 6, 4, 11, 11, 13, 13, 8, 12, 1, 3, 4, 7, 10, 14, 7,
		10, 9, 15, 5, 6, 0, 8, 15, 0, 14, 5, 2, 9, 3, 2, 12},
	{13, 1, 2, 15, 8, 13, 4, 8, 6, 10, 15, 3, 11, 7, 1, 4,
		10, 12, 9, 5, 3, 6, 14, 11, 5, 0, 0, 14, 12, 9, 7, 2,
		7, 2, 11, 1, 4, 14, 1, 7, 9, 4, 12, 10, 14, 8, 2, 13,
		0, 15, 6, 12, 10, 9, 13, 0, 15, 3, 3, 5, 5, 6, 8, 11},
}

var pTable = []int{
	16, 7, 20, 21, 29, 12, 28, 17,
	1, 15, 23, 26, 5, 18, 31, 10,
	2, 8, 24, 14, 32, 27, 3, 9,
	19, 13, 30, 6, 22, 11, 4, 25,
}

// Does not return a 1 for a 1 bit, it just returns non-zero
func getBit(array []byte, bit int) uint8 {
	location := bit / 8
	return uint8(array[location] & (0x80 >> (bit % 8)))
}

func setBit(array []byte, bit int) {
	location := bit / 8
	array[location] = array[location] | (0x80 >> (bit % 8))
}

func clearBit(array []byte, bit int) {
	location := bit / 8
	array[location] = array[location] & ^(0x80 >> (bit % 8))
}

func xor(target []byte, src []byte, length int) {
	for i := 0; i < length; i++ {
		target[i] = target[i] ^ src[i]
	}
}

func invertByte(b byte) byte {
	return 0xFF ^ b
}

func permute(target []byte, src []byte, permuteTable []int, length int) {
	for i := 0; i < length*8; i++ {
		if getBit(src, permuteTable[i]-1) != 0 {
			setBit(target, i)
		} else {
			clearBit(target, i)
		}
	}
}

func rol(target []byte) {
	carryLeft := (target[0] & 0x80) >> 3

	target[0] = (target[0] << 1) | ((target[1] & 0x80) >> 7)
	target[1] = (target[1] << 1) | ((target[2] & 0x80) >> 7)
	target[2] = (target[2] << 1) | ((target[3] & 0x80) >> 7)

	carryRight := (target[3] & 0x08) >> 3
	target[3] = (((target[3] << 1) | ((target[4] & 0x80) >> 7)) & 0xef) | carryLeft

	target[4] = (target[4] << 1) | ((target[5] & 0x80) >> 7)
	target[5] = (target[5] << 1) | ((target[6] & 0x80) >> 7)
	target[6] = (target[6] << 1) | carryRight
}

func ror(target []byte) {
	carryRight := (target[6] & 0x01) << 3

	target[6] = (target[6] >> 1) | ((target[5] & 0x01) << 7)
	target[5] = (target[5] >> 1) | ((target[4] & 0x01) << 7)
	target[4] = (target[4] >> 1) | ((target[3] & 0x01) << 7)

	carryLeft := (target[3] & 0x10) << 3
	target[3] = ((target[3]>>1)|
		((target[2]&0x01)<<7))&invertByte(0x08) | carryRight

	target[2] = (target[2] >> 1) | ((target[5] & 0x01) << 7)
	target[1] = (target[1] >> 1) | ((target[6] & 0x01) << 7)
	target[0] = (target[0] >> 1) | carryLeft
}

func memSet(a []byte, v byte) {
	for i := range a {
		a[i] = v
	}
}

func desBlockOperate(plaintext []byte, ciphertext []byte, key []byte, decrypt bool) {
	var ipBlock [DES_BLOCK_SIZE]byte
	var expansionBlock [EXPANSION_BLOCK_SIZE]byte
	var substitutionBlock [DES_BLOCK_SIZE / 2]byte
	var pboxTarget [DES_BLOCK_SIZE / 2]byte
	var recombBox [DES_BLOCK_SIZE / 2]byte

	var pc1Key [PC1_KEY_SIZE]byte
	var subKey [SUBKEY_SIZE]byte

	// Initial permutation
	permute(ipBlock[:], plaintext, ipTable, DES_BLOCK_SIZE)

	// Key schedule computation
	permute(pc1Key[:], key, pc1Table, PC1_KEY_SIZE)

	for round := 0; round < 16; round++ {
		// Feistel function on the first half of the block in ipBlock
		// Expansion - This permutation only looks at the first 4 bytes
		// (32 bits) of ipBlock; 16 of these are repeated in "expansionTable"
		permute(expansionBlock[:], ipBlock[4:], expansionTable, 6)

		// Key mixing
		if !decrypt {
			rol(pc1Key[:])
			if !(round <= 1 || round == 8 || round == 15) {
				rol(pc1Key[:])
			}
		}
		permute(subKey[:], pc1Key[:], pc2Table, SUBKEY_SIZE)

		if decrypt {
			ror(pc1Key[:])
			if !(round >= 14 || round == 7 || round == 0) {
				ror(pc1Key[:])
			}
		}

		xor(expansionBlock[:], subKey[:], SUBKEY_SIZE)
		memSet(substitutionBlock[:], 0)

		substitutionBlock[0] = byte(sbox[0][int(expansionBlock[0]&0xFC>>2)] << 4)
		substitutionBlock[0] |= byte(sbox[1][int((expansionBlock[0]&0x03)<<4|
			(expansionBlock[1]&0xF0)>>4)])

		substitutionBlock[1] = byte(sbox[2][int((expansionBlock[1]&0x0F)<<2|
			(expansionBlock[2]&0xC0)>>6)] << 4)
		substitutionBlock[1] |= byte(sbox[3][int(expansionBlock[2]&0x3F)])

		substitutionBlock[2] = byte(sbox[4][int((expansionBlock[3]&0xFC)>>2)] << 4)
		substitutionBlock[2] |= byte(sbox[5][int((expansionBlock[3]&0x03)<<4|
			(expansionBlock[4]&0xF0)>>4)])

		substitutionBlock[3] = byte(sbox[6][int((expansionBlock[4]&0x0F)<<2|
			(expansionBlock[5]&0xC0)>>6)] << 4)
		substitutionBlock[3] |= byte(sbox[7][int(expansionBlock[5]&0x3F)])

		permute(pboxTarget[:], substitutionBlock[:], pTable, DES_BLOCK_SIZE/2)

		copy(recombBox[:], ipBlock[:DES_BLOCK_SIZE/2])
		copy(ipBlock[:], ipBlock[4:])
		xor(recombBox[:], pboxTarget[:], DES_BLOCK_SIZE/2)
		copy(ipBlock[4:], recombBox[:])
	}

	copy(recombBox[:], ipBlock[:DES_BLOCK_SIZE/2])
	copy(ipBlock[:], ipBlock[4:])
	copy(ipBlock[4:], recombBox[:])

	permute(ciphertext, ipBlock[:], fpTable, DES_BLOCK_SIZE)
}

func desOperate(input []byte, output []byte, iv []byte, key []byte, decrypt bool) {
	if len(input)%DES_BLOCK_SIZE != 0 {
		log.Fatal("Input byte size needs to be a multiple of 8")
	}

	var inputBlock [DES_BLOCK_SIZE]byte

	inputLen := len(input)
	inputStart := 0
	outputStart := 0
	for inputLen != 0 {
		copy(inputBlock[:], input[inputStart:(inputStart+DES_BLOCK_SIZE)])
		if !decrypt {
			xor(inputBlock[:], iv, DES_BLOCK_SIZE)
			desBlockOperate(inputBlock[:], output[outputStart:(outputStart+DES_BLOCK_SIZE)], key, decrypt)
			copy(iv, output[outputStart:(outputStart+DES_BLOCK_SIZE)])
		} else {
			desBlockOperate(inputBlock[:], output[outputStart:(outputStart+DES_BLOCK_SIZE)], key, decrypt)
			xor(output[outputStart:(outputStart+DES_BLOCK_SIZE)], iv, DES_BLOCK_SIZE)
			copy(iv[:], input[inputStart:(inputStart+DES_BLOCK_SIZE)])

		}
		inputStart += DES_BLOCK_SIZE
		outputStart += DES_BLOCK_SIZE
		inputLen -= DES_BLOCK_SIZE
	}
}

func desDecrypt(plaintext []byte, ciphertext []byte, iv []byte, key []byte) {
	desOperate(plaintext, iv, ciphertext, key, true)
}

func DesEncrypt(plaintext []byte, ciphertext []byte, iv []byte, key []byte) {
	desOperate(plaintext, ciphertext, iv, key, false)
}
