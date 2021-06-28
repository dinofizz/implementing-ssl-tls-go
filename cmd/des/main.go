package main

import (
	"fmt"
	"github.com/dinofizz/impl-tsl-go/pkg/common"
	"github.com/dinofizz/impl-tsl-go/pkg/des"
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

	triplicate := false
	if len(key) == 24 {
		triplicate = true
	}

	if args[1] == "-e" {
		output := des.Encrypt(input, iv, key, triplicate)
		fmt.Println(common.HexDisplay(output))
	} else if args[1] == "-d" {
		output := des.Decrypt(input, iv, key, triplicate)
		fmt.Println(common.HexDisplay(output))
	} else {
		log.Fatalf("Usage: %v [-e|-d] <key> <iv> <input>\n", args[0])
	}
}