package des

import (
	"github.com/dinofizz/impl-tsl-go/pkg/common"
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
			common.Decode([]byte("0x71828547387b18e5")),
			false,
		},
		{
			"triple des encrypt",
			[]byte("abcdefgh"),
			[]byte("initialz"),
			[]byte("twentyfourcharacterinput"),
			common.Decode([]byte("0xc0c48bc47e87ce17")),
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
			common.Decode([]byte("0x71828547387b18e5")),
			[]byte("initialz"),
			[]byte("password"),
			common.Decode([]byte("abcdefgh")),
			false,
		},
		{
			"triple des decrypt",
			common.Decode([]byte("0xc0c48bc47e87ce17")),
			[]byte("initialz"),
			[]byte("twentyfourcharacterinput"),
			common.Decode([]byte("abcdefgh")),
			true,
		},
	}

	for _, tt := range destests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, len(tt.input))
			DesDecrypt(tt.input, output, tt.iv, tt.key, tt.triplicate)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
