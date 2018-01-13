package main

import (
	"fmt"
	"os"

	"github.com/suspectpart/lsgpt/gpt"
)

const (
	errorCode = 1
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("lsgpt: missing file operand")
		os.Exit(errorCode)
	}

	filename := os.Args[1]

	table, err := gpt.ReadFrom(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(errorCode)
	}

	fmt.Println(table)
}
