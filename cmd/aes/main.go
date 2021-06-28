package main

import (
	"fmt"
	"github.com/dinofizz/impl-tsl-go/pkg/aes"
	"github.com/dinofizz/impl-tsl-go/pkg/common"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		log.Fatalf("Usage: %v [-e|-d] <key> <iv> <input>\n", args[0])
	}

	iv := make([]byte, 16,16)

	key := []byte(args[2])
	ivBytes := []byte(args[3])
	copy(iv, ivBytes)
	input := []byte(args[4])

	key = common.Decode(key)
	keyLen := len(key)

	if keyLen != 16 && keyLen != 32 {
		log.Fatalf("Unsupported key length: %d", keyLen)
	}

	iv = common.Decode(iv)
	input = common.Decode(input)

	w := make([][]byte, 60,60)

	for i := 0; i < 60; i++ {
		v := make([]byte, 4, 4)
		w[i] = v
	}

	aes.ComputeKeySchedule(key, w)
	if args[1] == "-e" {
		output := aes.Encrypt(input, iv, key)
		fmt.Println(common.HexDisplay(output))
	} else if args[1] == "-d" {
		output :=aes.Decrypt(input, iv, key)
		fmt.Println(common.HexDisplay(output))
	} else {
		log.Fatalf("Usage: %v [-e|-d] <key> <iv> <input>\n", args[0])
	}
}