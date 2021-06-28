package rc4

import (
	"github.com/dinofizz/impl-tsl-go/pkg/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encrypt(t *testing.T) {
	var aestests = []struct {
		name           string
		input          []byte
		key            []byte
		expectedOutput []byte
	}{
		{
			"rc4 encrypt",
			[]byte("abcdefghijklmnopqrstuvqxyz1234567890ABCDEFGHJIKL"),
			[]byte("passwordPASSWORD"),
			common.Decode([]byte("0x861eb964f550d95c6961cb2978d0263e728baf8ee05d2d5a7feb587c897c18ca384cc72db9d9a9c7e6364ec6a95dbe62")),
		},
	}

	for _, tt := range aestests {
		t.Run(tt.name, func(t *testing.T) {
			state := Initialize()
			output := Operate(tt.input, tt.key, state)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_decrypt(t *testing.T) {
	var aestests = []struct {
		name           string
		input          []byte
		key            []byte
		expectedOutput []byte
		triplicate     bool
	}{
		{
			"aes decrypt",
			common.Decode([]byte("0x861eb964f550d95c6961cb2978d0263e728baf8ee05d2d5a7feb587c897c18ca384cc72db9d9a9c7e6364ec6a95dbe62")),
			[]byte("passwordPASSWORD"),
			common.Decode([]byte("abcdefghijklmnopqrstuvqxyz1234567890ABCDEFGHJIKL")),
			false,
		},
	}

	for _, tt := range aestests {
		t.Run(tt.name, func(t *testing.T) {
			state := Initialize()
			output := Operate(tt.input, tt.key, state)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
