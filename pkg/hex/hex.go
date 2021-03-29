package hex

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func Decode(input []byte) (decoded []byte) {
	s := string(input[:2])
	if s != "0x" {
		return input
	}
	d := make([]byte, (len(input) >> 1) -1)
	for i := 2; i < len(input); i += 2 {
		var x, y int
		if int(input[i]) <= int('9') {
			x = int(input[i]) - int('0')
		} else {
			r, _ := utf8.DecodeRune(input[i:])
			rl := unicode.ToLower(r)
			x = int(rl) - int('a') + 10
		}
		x = x << 4

		if int(input[i+1]) <= int('9') {
			y = int(input[i+1]) - int('0')
		} else {
			r, _ := utf8.DecodeRune(input[i+1:])
			rl := unicode.ToLower(r)
			y = int(rl) - int('a') + 10
		}
		d[(i/2)-1] = byte(x | y)
	}

	return d
}

func HexDisplay(h []byte) string {
	var result string
	for _, c := range h {
		result += fmt.Sprintf("%.02x", c)
	}
	return result
}