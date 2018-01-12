package main

import (
	"fmt"
	"os"

	"github.com/suspectpart/lsgpt/gpt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file / drive specified")
		os.Exit(1)
	}

	gptHeader, err := gpt.Read(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Found GPT Header")
	fmt.Printf("Disk UUID: %s\n", gptHeader.DiskGUID.AsUUID())
	fmt.Printf("First Usable LBA (usually 34): %d\n", gptHeader.FirstUsableLBA)
	fmt.Printf("Last Usable LBA: %d\n", gptHeader.LastUsableLBA)
	fmt.Printf("Starting LBA of array of partition entries (always 2 in primary copy): %d\n", gptHeader.StartLBA)
}
