package hex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decode(t *testing.T) {
	in := []byte("0x71828547387b18e5")
	out := Decode(in)
	assert.Equal(t, "6162636465666768", out)
}
