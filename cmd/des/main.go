package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dinofizz/impl-tsl-go/pkg/des"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		log.Fatalf("Usage: %v <key> <iv> <input<\n", args[0])
	}

	key := []byte(args[1])
	iv := []byte(args[2])
	input := []byte(args[3])
	output := make([]byte, len(input))

	des.DesEncrypt(input, output, iv, key)
	for _, c := range output {
		fmt.Printf("%.02x", c)
	}
	fmt.Println()
}