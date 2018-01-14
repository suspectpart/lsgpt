package main

import (
	"fmt"
	"os"

	"github.com/suspectpart/lsgpt/gpt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("lsgpt: missing file operand")
		os.Exit(1)
	}

	filename := os.Args[1]

	table, err := gpt.ReadFrom(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(table)
}
