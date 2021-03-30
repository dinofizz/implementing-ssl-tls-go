package des

import (
	"github.com/dinofizz/impl-tsl-go/pkg/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encrypt(t *testing.T) {
	var destests = []struct {
		name           string
		input          []byte
		iv             []byte
		key            []byte
		expectedOutput []byte
		triplicate     bool
	}{
		{
			"des encrypt",
			[]byte("abcdefgh"),
			[]byte("initialz"),
			[]byte("password"),
			[]byte{0x71, 0x82, 0x85, 0x47, 0x38, 0x7b, 0x18, 0xe5},
			false,
		},
		{
			"triple des encrypt",
			[]byte("abcdefgh"),
			[]byte("initialz"),
			[]byte("twentyfourcharacterinput"),
			[]byte{0xc0, 0xc4, 0x8b, 0xc4, 0x7e, 0x87, 0xce, 0x17},
			true,
		},
	}

	for _, tt := range destests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, len(tt.input))
			DesEncrypt(tt.input, output, tt.iv, tt.key, tt.triplicate)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_decrypt(t *testing.T) {
	var destests = []struct {
		name           string
		input          []byte
		iv             []byte
		key            []byte
		expectedOutput []byte
		triplicate     bool
	}{
		{
			"des decrypt",
			[]byte("0x71828547387b18e5"),
			[]byte("initialz"),
			[]byte("password"),
			[]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68},
			false,
		},
		{
			"triple des decrypt",
			[]byte("0xc0c48bc47e87ce17"),
			[]byte("initialz"),
			[]byte("twentyfourcharacterinput"),
			[]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68},
			true,
		},
	}

	for _, tt := range destests {
		t.Run(tt.name, func(t *testing.T) {
			input := hex.Decode(tt.input)
			output := make([]byte, len(input))
			DesDecrypt(input, output, tt.iv, tt.key, tt.triplicate)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
