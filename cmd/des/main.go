package main

import (
	"fmt"
	"github.com/dinofizz/impl-tsl-go/pkg/des"
	"github.com/dinofizz/impl-tsl-go/pkg/hex"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		log.Fatalf("Usage: %v [-e|-d] <key> <iv> <input>\n", args[0])
	}


	key := []byte(args[2])
	iv := []byte(args[3])
	input := []byte(args[4])

	key = hex.Decode(key)
	iv = hex.Decode(iv)
	input = hex.Decode(input)

	output := make([]byte, len(input))

	if args[1] == "-e" {
		des.DesEncrypt(input, output, iv, key)
		fmt.Println(hex.HexDisplay(output))
	} else if args[1] == "-d" {
		des.DesDecrypt(input, output, iv, key)
		fmt.Println(hex.HexDisplay(output))
	} else {
		log.Fatalf("Usage: %v [-e|-d] <key> <iv> <input>\n", args[0])
	}
}