package main

import (
	"fmt"
	"github.com/dinofizz/impl-tsl-go/pkg/common"
	"github.com/dinofizz/impl-tsl-go/pkg/rc4"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		log.Fatalf("Usage: %v [-e|-d] <key> <input>\n", args[0])
	}

	key := []byte(args[2])
	input := []byte(args[3])

	key = common.Decode(key)
	input = common.Decode(input)

	state := rc4.Initialize()
	output := rc4.Operate(input, key, state)
	fmt.Println(common.HexDisplay(output))
}
