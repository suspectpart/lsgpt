package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/suspectpart/lsgpt/gpt"
)

// constants
const (
	ExitNoDriveSpecified = 1
	ExitFileNotFound     = 2
	ExitNoGPT            = 3
	GPTHeaderStart       = 512
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file / drive specified")
		os.Exit(ExitNoDriveSpecified)
	}

	drive, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("File not found")
		os.Exit(ExitFileNotFound)
	}

	defer drive.Close()

	gptHeader := make([]byte, 512)

	drive.Seek(GPTHeaderStart, 0)
	drive.Read(gptHeader)

	buffer := bytes.NewBuffer(gptHeader)

	header := gpt.Header{}

	err = binary.Read(buffer, binary.LittleEndian, &header)

	if string(header.Signature[:8]) != "EFI PART" {
		fmt.Println("No EFI header found")
		os.Exit(ExitNoGPT)
	}

	fmt.Println("Found GPT Header")
	fmt.Printf("Disk UUID: %s\n", header.DiskGUID.AsUUID())
	fmt.Printf("First Usable LBA (usually 34): %d\n", header.FirstUsableLBA)
	fmt.Printf("Last Usable LBA: %d\n", header.LastUsableLBA)
	fmt.Printf("Starting LBA of array of partition entries (always 2 in primary copy): %d\n", header.StartLBA)
}
