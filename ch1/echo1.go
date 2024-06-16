package ch1

import (
	"fmt"
	"os"
	"strings"
)

func GetOsArgs() {
	var end string
	args := strings.Split(os.Args[0], "/")
	end = args[len(args)-1]
	fmt.Println(end)
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
