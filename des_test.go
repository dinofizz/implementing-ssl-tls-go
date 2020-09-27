package tlsimpl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getBit(t *testing.T) {
	var bitTests = []struct {
		a   []byte
		n   uint8 //
		set bool  // expected result
	}{
		{[]byte{0x55}, 0, false},
		{[]byte{0x55}, 1, true},
		{[]byte{0x55}, 2, false},
		{[]byte{0x55}, 3, true},
		{[]byte{0x55}, 4, false},
		{[]byte{0x55}, 5, true},
		{[]byte{0x55}, 6, false},
		{[]byte{0x55}, 7, true},
		{[]byte{0xaa, 0x55, 0xaa}, 16, true},
		{[]byte{0xaa, 0x55, 0xaa}, 17, false},
		{[]byte{0xaa, 0x55, 0xaa}, 18, true},
		{[]byte{0xaa, 0x55, 0xaa}, 19, false},
		{[]byte{0xaa, 0x55, 0xaa}, 20, true},
		{[]byte{0xaa, 0x55, 0xaa}, 21, false},
		{[]byte{0xaa, 0x55, 0xaa}, 22, true},
		{[]byte{0xaa, 0x55, 0xaa}, 23, false},
	}
	for _, tt := range bitTests {
		bit := getBit(tt.a, tt.n)
		if tt.set {
			assert.NotEqual(t, uint8(0), bit)
		} else {
			assert.Equal(t, uint8(0), bit)
		}
	}
}

func Test_setBit(t *testing.T) {
	var bitTests = []struct {
		a      []byte
		n      uint8  //
		result []byte // expected result
	}{
		{[]byte{0x00}, 0, []byte{0x80}},
		{[]byte{0x00}, 1, []byte{0x40}},
		{[]byte{0x00}, 2, []byte{0x20}},
		{[]byte{0x00}, 3, []byte{0x10}},
		{[]byte{0x00}, 4, []byte{0x08}},
		{[]byte{0x00}, 5, []byte{0x04}},
		{[]byte{0x00}, 6, []byte{0x02}},
		{[]byte{0x00}, 7, []byte{0x01}},
		{[]byte{0xFF}, 0, []byte{0xFF}},
		{[]byte{0xFF}, 1, []byte{0xFF}},
		{[]byte{0xFF}, 2, []byte{0xFF}},
		{[]byte{0xFF}, 3, []byte{0xFF}},
		{[]byte{0xFF}, 4, []byte{0xFF}},
		{[]byte{0xFF}, 5, []byte{0xFF}},
		{[]byte{0xFF}, 6, []byte{0xFF}},
		{[]byte{0xFF}, 7, []byte{0xFF}},
		{[]byte{0xaa, 0x55, 0x00}, 16, []byte{0xaa, 0x55, 0x80}},
		{[]byte{0xaa, 0x55, 0x00}, 17, []byte{0xaa, 0x55, 0x40}},
		{[]byte{0xaa, 0x55, 0x00}, 18, []byte{0xaa, 0x55, 0x20}},
		{[]byte{0xaa, 0x55, 0x00}, 19, []byte{0xaa, 0x55, 0x10}},
		{[]byte{0xaa, 0x55, 0x00}, 20, []byte{0xaa, 0x55, 0x08}},
		{[]byte{0xaa, 0x55, 0x00}, 21, []byte{0xaa, 0x55, 0x04}},
		{[]byte{0xaa, 0x55, 0x00}, 22, []byte{0xaa, 0x55, 0x02}},
		{[]byte{0xaa, 0x55, 0x00}, 23, []byte{0xaa, 0x55, 0x01}},
	}
	for _, tt := range bitTests {
		setBit(&tt.a, tt.n)
		assert.Equal(t, tt.a, tt.result)
	}

}
func Test_clearBit(t *testing.T) {
	var bitTests = []struct {
		a      []byte
		n      uint8  //
		result []byte // expected result
	}{
		{[]byte{0xFF}, 0, []byte{0x7f}},
		{[]byte{0xFF}, 1, []byte{0xbf}},
		{[]byte{0xFF}, 2, []byte{0xdf}},
		{[]byte{0xFF}, 3, []byte{0xef}},
		{[]byte{0xFF}, 4, []byte{0xf7}},
		{[]byte{0xFF}, 5, []byte{0xfb}},
		{[]byte{0xFF}, 6, []byte{0xfd}},
		{[]byte{0xFF}, 7, []byte{0xfe}},
		{[]byte{0x00}, 0, []byte{0x00}},
		{[]byte{0x00}, 1, []byte{0x00}},
		{[]byte{0x00}, 2, []byte{0x00}},
		{[]byte{0x00}, 3, []byte{0x00}},
		{[]byte{0x00}, 4, []byte{0x00}},
		{[]byte{0x00}, 5, []byte{0x00}},
		{[]byte{0x00}, 6, []byte{0x00}},
		{[]byte{0x00}, 7, []byte{0x00}},
		{[]byte{0xaa, 0x55, 0xFF}, 16, []byte{0xaa, 0x55, 0x7f}},
		{[]byte{0xaa, 0x55, 0xFF}, 17, []byte{0xaa, 0x55, 0xbf}},
		{[]byte{0xaa, 0x55, 0xFF}, 18, []byte{0xaa, 0x55, 0xdf}},
		{[]byte{0xaa, 0x55, 0xFF}, 19, []byte{0xaa, 0x55, 0xef}},
		{[]byte{0xaa, 0x55, 0xFF}, 20, []byte{0xaa, 0x55, 0xf7}},
		{[]byte{0xaa, 0x55, 0xFF}, 21, []byte{0xaa, 0x55, 0xfb}},
		{[]byte{0xaa, 0x55, 0xFF}, 22, []byte{0xaa, 0x55, 0xfd}},
		{[]byte{0xaa, 0x55, 0xFF}, 23, []byte{0xaa, 0x55, 0xfe}},
	}
	for _, tt := range bitTests {
		clearBit(&tt.a, tt.n)
		assert.Equal(t, tt.a, tt.result)
	}
}

func Test_xor(t *testing.T) {
	var xorTests = []struct {
		target []byte
		src []byte
	}{
		{[]byte{0x55}, []byte{0xaa}},
		{[]byte{0x00}, []byte{0xFF}},
		{[]byte{0x7c}, []byte{0xe5}},
		{[]byte{0x37}, []byte{0x88}},
	}

	for _,tt := range xorTests {
		orig := make([]byte, len(tt.target))
		copy(orig, tt.target)
		xor(&tt.target, &tt.src)
		xor(&tt.target, &tt.src)
		assert.Equal(t, orig, tt.target)
	}
}
