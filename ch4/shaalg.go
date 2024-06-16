package ch4

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var alg = flag.Int("a", 256, "256 by default or pass 224 or 512")

func Shaalg() {
	input := bufio.NewScanner(os.Stdin)
	flag.Parse()

	if *alg != 256 && *alg != 384 && *alg != 512 {
		*alg = 256
	}
	for input.Scan() {
		text := input.Text()
		if *alg == 256 {
			fmt.Printf("256: %x\n", sha256.Sum256([]byte(text)))
		} else if *alg == 384 {
			fmt.Printf("384: %x\n", sha512.Sum384([]byte(text)))
		} else {
			fmt.Printf("512: %x\n", sha512.Sum512([]byte(text)))
		}
	}

}
