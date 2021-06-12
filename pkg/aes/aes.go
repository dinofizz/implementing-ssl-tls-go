package aes

import "github.com/dinofizz/impl-tsl-go/pkg/common"

func rotWord(w []byte) {
	var tmp byte
	tmp = w[0]
	w[0] = w[1]
	w[1] = w[2]
	w[2] = w[3]
	w[3] = tmp
}

var sbox = [][]byte{
	{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5,
		0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
	{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0,
		0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
	{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc,
		0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
	{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a,
		0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
	{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0,
		0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
	{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b,
		0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
	{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85,
		0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
	{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5,
		0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
	{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17,
		0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
	{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88,
		0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
	{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c,
		0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
	{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9,
		0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
	{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6,
		0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
	{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e,
		0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
	{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94,
		0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
	{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68,
		0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
}

var invSbox = [][]byte{
	{0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38,
		0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb},
	{0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87,
		0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb},
	{0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d,
		0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e},
	{0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2,
		0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25},
	{0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16,
		0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92},
	{0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda,
		0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84},
	{0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a,
		0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06},
	{0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02,
		0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b},
	{0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea,
		0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73},
	{0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85,
		0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e},
	{0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89,
		0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b},
	{0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20,
		0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4},
	{0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31,
		0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f},
	{0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d,
		0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef},
	{0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0,
		0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61},
	{0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26,
		0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d},
}

func subWord(w []byte) {
	for i := 0; i < 4; i++ {
		w[i] = sbox[int(w[i]&0xf0)>>4][int(w[i]&0x0f)]
	}
}

func ComputeKeySchedule(key []byte, w [][]byte) {
	keyWords := len(key) >> 2

	rcon := byte(0x01)

	copy(w[0], key[0:4])
	copy(w[1], key[4:8])
	copy(w[2], key[8:12])
	copy(w[3], key[12:16])

	for i := keyWords; i < 4*(keyWords+7); i++ {
		copy(w[i], w[i-1])

		if i%keyWords == 0 {
			rotWord(w[i])
			subWord(w[i])
			if i%36 == 0 {
				rcon = byte(0x1b)
			}
			w[i][0] ^= rcon
			rcon <<= 1
		} else if keyWords > 6 && i%keyWords == 4 {
			subWord(w[i])
		}

		w[i][0] = w[i][0] ^ w[i-keyWords][0]
		w[i][1] ^= w[i-keyWords][1]
		w[i][2] ^= w[i-keyWords][2]
		w[i][3] ^= w[i-keyWords][3]
	}
}

func addRoundKey(state [][]byte, w [][]byte) {
	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			state[r][c] = state[r][c] ^ w[c][r]
		}
	}
}

func subBytes(state [][]byte) {
	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			state[r][c] = sbox[(state[r][c]&0xF0)>>4][state[r][c]&0x0F]
		}
	}
}

func invSubBytes(state [][]byte) {
	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			state[r][c] = invSbox[(state[r][c]&0xF0)>>4][state[r][c]&0x0F]
		}
	}
}

func shiftRows(state [][]byte) {
	tmp := state[1][0]
	state[1][0] = state[1][1]
	state[1][1] = state[1][2]
	state[1][2] = state[1][3]
	state[1][3] = tmp

	tmp = state[2][0]
	state[2][0] = state[2][2]
	state[2][2] = tmp
	tmp = state[2][1]
	state[2][1] = state[2][3]
	state[2][3] = tmp

	tmp = state[3][3]
	state[3][3] = state[3][2]
	state[3][2] = state[3][1]
	state[3][1] = state[3][0]
	state[3][0] = tmp
}

func invShiftRows(state [][]byte) {
	tmp := state[1][2]
	state[1][2] = state[1][1]
	state[1][1] = state[1][0]
	state[1][0] = state[1][3]
	state[1][3] = tmp

	tmp = state[2][0]
	state[2][0] = state[2][2]
	state[2][2] = tmp
	tmp = state[2][1]
	state[2][1] = state[2][3]
	state[2][3] = tmp

	tmp = state[3][0]
	state[3][0] = state[3][1]
	state[3][1] = state[3][2]
	state[3][2] = state[3][3]
	state[3][3] = tmp
}

func xtime(x byte) byte {
	var m byte
	if x&0x80 != 0x00 {
		m = 0x1b
	} else {
		m = 0x00
	}
	return (x << 1) ^ m
}

func dot(x byte, y byte) byte {
	var product byte
	for mask := byte(0x01); mask != 0x00; mask <<= 1 {
		if y&mask != 0 {
			product ^= x
		}
		x = xtime(x)
	}
	return product
}

func mixColumns(s [][]byte) {
	var t [4]byte
	for c := 0; c < 4; c++ {
		t[0] = dot(2, s[0][c]) ^ dot(3, s[1][c]) ^ s[2][c] ^ s[3][c]
		t[1] = s[0][c] ^ dot(2, s[1][c]) ^ dot(3, s[2][c]) ^ s[3][c]
		t[2] = s[0][c] ^ s[1][c] ^ dot(2, s[2][c]) ^ dot(3, s[3][c])
		t[3] = dot(3, s[0][c]) ^ s[1][c] ^ s[2][c] ^ dot(2, s[3][c])
		s[0][c] = t[0]
		s[1][c] = t[1]
		s[2][c] = t[2]
		s[3][c] = t[3]
	}
}

func invMixColumns(s [][]byte) {
	var t [4]byte
	for c := 0; c < 4; c++ {
		t[0] = dot(0x0e, s[0][c]) ^ dot(0x0b, s[1][c]) ^
			dot(0x0d, s[2][c]) ^ dot(0x09, s[3][c])
		t[1] = dot(0x09, s[0][c]) ^ dot(0x0e, s[1][c]) ^
			dot(0x0b, s[2][c]) ^ dot(0x0d, s[3][c])
		t[2] = dot(0x0d, s[0][c]) ^ dot(0x09, s[1][c]) ^
			dot(0x0e, s[2][c]) ^ dot(0x0b, s[3][c])
		t[3] = dot(0x0b, s[0][c]) ^ dot(0x0d, s[1][c]) ^
			dot(0x09, s[2][c]) ^ dot(0x0e, s[3][c])
		s[0][c] = t[0]
		s[1][c] = t[1]
		s[2][c] = t[2]
		s[3][c] = t[3]
	}
}

func createSliceOfSlices(c, r int) [][]byte {
	s := make([][]byte, c, c)
	for i := 0; i < c; i++ {
		s[i] = make([]byte, r, r)
	}
	return s
}

func aesBlockEncrypt(inputBlock []byte, outputBlock []byte, outputIndex int, key []byte) {
	state := createSliceOfSlices(4, 4)
	w := createSliceOfSlices(60, 4)

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			state[r][c] = inputBlock[r+(4*c)]
		}
	}

	//rounds = key size in 4 byte words + 6
	nr := (len(key) >> 2) + 6
	ComputeKeySchedule(key, w)

	addRoundKey(state, w)

	for round := 0; round < nr; round++ {
		subBytes(state)
		shiftRows(state)
		if round < (nr - 1) {
			mixColumns(state)
		}
		addRoundKey(state, w[(round+1)*4:])
	}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			outputBlock[outputIndex+(r+(4*c))] = state[r][c]
		}
	}
}

func aesBlockDecrypt(inputBlock []byte, inputIndex int, outputBlock []byte, outputIndex int, key []byte) {
	state := createSliceOfSlices(4, 4)
	w := createSliceOfSlices(60, 4)

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			state[r][c] = inputBlock[inputIndex + (r+(4*c))]
		}
	}

	//rounds = key size in 4 byte words + 6
	nr := (len(key) >> 2) + 6
	ComputeKeySchedule(key, w)

	addRoundKey(state, w[nr*4:])

	for round := nr; round > 0; round-- {
		invShiftRows(state)
		invSubBytes(state)
		addRoundKey(state, w[(round-1)*4:])
		if round > 1 {
			invMixColumns(state)
		}
	}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			outputBlock[outputIndex + (r+(4*c))] = state[r][c]
		}
	}
}

const AES_BLOCK_SIZE = 16

func AesEncrypt(input, output, iv, key []byte) {
	inputBlock := make([]byte, AES_BLOCK_SIZE, AES_BLOCK_SIZE)

	inputLen := len(input)
	inputIndex := 0
	outputIndex := 0
	for ; inputLen >= AES_BLOCK_SIZE; {
		copy(inputBlock, input[inputIndex:inputIndex+AES_BLOCK_SIZE])
		common.Xor(inputBlock, iv, AES_BLOCK_SIZE)
		aesBlockEncrypt(inputBlock, output, outputIndex, key)
		copy(iv, output[outputIndex:outputIndex+AES_BLOCK_SIZE])
		inputIndex += AES_BLOCK_SIZE
		outputIndex += AES_BLOCK_SIZE
		inputLen -= AES_BLOCK_SIZE
	}
}

func AesDecrypt(input, output, iv, key []byte) {
	inputLen := len(input)
	inputIndex := 0
	outputIndex := 0
	for ; inputLen >= AES_BLOCK_SIZE; {
		aesBlockDecrypt(input, inputIndex, output, outputIndex, key)
		common.Xor(output[outputIndex:outputIndex+AES_BLOCK_SIZE], iv, AES_BLOCK_SIZE)
		copy(iv, input[inputIndex:inputIndex+AES_BLOCK_SIZE])
		inputIndex += AES_BLOCK_SIZE
		outputIndex += AES_BLOCK_SIZE
		inputLen -= AES_BLOCK_SIZE
	}
}
