package tlsimpl

import (
	"github.com/stretchr/testify/assert"
	"github.com/dinofizz/impl-tsl-go/pkg/des"
	"testing"
)

func Test_encrypt(t *testing.T) {
	input := []byte("abcdefgh")
	iv := []byte("initialz")
	key := []byte("password")
	output := make([]byte, len(input))
	expectedOutput := []byte{0x71, 0x82, 0x85, 0x47, 0x38, 0x7b, 0x18, 0xe5}
	des.DesEncrypt(input, output, iv, key)
	assert.Equal(t, expectedOutput, output, )
}