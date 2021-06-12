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


	key := []byte(args[2])
	iv := []byte(args[3])
	input := []byte(args[4])

	key = common.Decode(key)
	iv = common.Decode(iv)
	input = common.Decode(input)

	w := make([][]byte, 60,60)

	for i := 0; i < 60; i++ {
		v := make([]byte, 4, 4)
		w[i] = v
	}

	output := make([]byte, len(input))

	aes.ComputeKeySchedule(key, w)
	if args[1] == "-e" {
		aes.AesEncrypt(input, output, iv, key)
		fmt.Println(common.HexDisplay(output))
	} else if args[1] == "-d" {
		aes.AesDecrypt(input, output, iv, key)
		fmt.Println(common.HexDisplay(output))
	} else {
		log.Fatalf("Usage: %v [-e|-d] <key> <iv> <input>\n", args[0])
	}
}