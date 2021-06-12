package aes

import (
	"github.com/dinofizz/impl-tsl-go/pkg/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encrypt(t *testing.T) {
	var aestests = []struct {
		name           string
		input          []byte
		iv             []byte
		key            []byte
		expectedOutput []byte
	}{
		{
			"aes encrypt",
			[]byte("abcdefghijklmnopqrstuvqxyz1234567890ABCDEFGHJIKL"),
			[]byte("initialzINITIALZ"),
			[]byte("passwordPASSWORD"),
			common.Decode([]byte("0xbe24b7b8b38849fbe1fb621f7e390a6e32a465ea0a658c2cfc914e28d0b81b48462e2ce2e1f7ad88fe28cd8cb798c838")),
		},
	}

	for _, tt := range aestests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, len(tt.input))
			AesEncrypt(tt.input, output, tt.iv, tt.key)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_decrypt(t *testing.T) {
	var aestests = []struct {
		name           string
		input          []byte
		iv             []byte
		key            []byte
		expectedOutput []byte
		triplicate     bool
	}{
		{
			"aes decrypt",
			common.Decode([]byte("0xbe24b7b8b38849fbe1fb621f7e390a6e32a465ea0a658c2cfc914e28d0b81b48462e2ce2e1f7ad88fe28cd8cb798c838")),
			[]byte("initialzINITIALZ"),
			[]byte("passwordPASSWORD"),
			common.Decode([]byte("abcdefghijklmnopqrstuvqxyz1234567890ABCDEFGHJIKL")),
			false,
		},
	}

	for _, tt := range aestests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, len(tt.input))
			AesDecrypt(tt.input, output, tt.iv, tt.key)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
