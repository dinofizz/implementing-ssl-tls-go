package common

import (
	"errors"
	"fmt"
)

var base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var unbase64 = [...]byte{
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 62, 80, 80, 80, 63,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 80, 80, 80, 64, 80, 80,
	80, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 80, 80, 80, 80, 80,
	80, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 80, 80, 80, 80, 80}

func Base64Decode(input *[]byte) (*[]byte, error) {
	inputLen := len(*input)
	output := make([]byte, 0, inputLen)

	if inputLen&0x03 != 0 {
		return nil, errors.New("Invalid input buffer length")
	}

	for inputIndex := 0; ; inputLen -= 4 {
		if inputLen == 0 {
			break
		}

		for i := 0; i <= 3; i++ {
			if (*input)[i] > 128 || unbase64[(*input)[i]] == byte(int(80)) {
				return nil, errors.New(fmt.Sprintf("Invalid character for base64 encoding: %c", (*input)[i]))
			}
		}
		val := unbase64[(*input)[inputIndex]]<<2 | (unbase64[(*input)[inputIndex+1]]&0x30)>>4
		output = append(output, val)

		if (*input)[inputIndex+2] != []byte("=")[0] {
			val = (unbase64[(*input)[inputIndex+1]]&0x0F)<<4 | (unbase64[(*input)[inputIndex+2]]&0x3C)>>2
			output = append(output, val)
		}

		if (*input)[inputIndex+3] != []byte("=")[0] {
			val = (unbase64[(*input)[inputIndex+2]]&0x03)<<6 | unbase64[(*input)[inputIndex+3]]
			output = append(output, val)
		}

		inputIndex += 4
	}

	return &output, nil
}

func Base64Encode(input *[]byte) *[]byte {
	if len(*input) == 0 {
		output := make([]byte, 0, 0)
		return &output
	}

	inputLen := len(*input)
	outputLen := inputLen * 4 / 3
	output := make([]byte, 0, outputLen)
	for i := 0; ; inputLen -= 3 {
		if inputLen == 0 {
			break
		}

		index := ((*input)[i] & 0xFC) >> 2
		output = append(output, base64[index])

		if inputLen == 1 {
			index = ((*input)[i] & 0x03) << 4
			output = append(output, base64[index])
			output = append(output, []byte("==")...)
			break
		}

		index = (((*input)[i] & 0x03) << 4) | (((*input)[i+1] & 0xF0) >> 4)

		output = append(output, base64[index])

		if inputLen == 2 {
			index = ((*input)[i+1] & 0x0F) << 2
			output = append(output, base64[index])
			output = append(output, []byte("=")...)
			break
		}

		index = (((*input)[i+1] & 0x0F) << 2) | (((*input)[i+2] & 0xC0) >> 6)
		output = append(output, base64[index])

		index = (*input)[i+2] & 0x3F
		output = append(output, base64[index])
		i += 3
	}

	return &output
}
