package des

import (
	"github.com/dinofizz/impl-tsl-go/pkg/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encrypt(t *testing.T) {
	input := []byte("abcdefgh")
	iv := []byte("initialz")
	key := []byte("password")
	output := make([]byte, len(input))
	expectedOutput := []byte{0x71, 0x82, 0x85, 0x47, 0x38, 0x7b, 0x18, 0xe5}
	DesEncrypt(input, output, iv, key)
	assert.Equal(t, expectedOutput, output, )
}

func Test_decrypt(t *testing.T) {
	input := []byte("0x71828547387b18e5")
	input = hex.Decode(input)
	iv := []byte("initialz")
	key := []byte("password")
	output := make([]byte, len(input))
	expectedOutput := "6162636465666768"
	DesDecrypt(input, output, iv, key)
	s := hex.HexDisplay(output)
	assert.Equal(t, expectedOutput, s, )
}