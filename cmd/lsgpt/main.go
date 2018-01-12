package main

import (
	"fmt"
	"os"

	"github.com/suspectpart/lsgpt/gpt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file specified")
		os.Exit(1)
	}

	header, err := gpt.ReadHeader(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("=== GPT Header ===")
	fmt.Printf("Disk UUID:\t\t\t%s\n", header.DiskGUID.AsUUID())
	fmt.Printf("Header Checksum CRC32:\t\t%d\n", header.HeaderCRC32)
	fmt.Printf("Header Size:\t\t\t%d\n", header.HeaderSize)
	fmt.Printf("First Usable LBA:\t\t%d\n", header.FirstUsableLBA)
	fmt.Printf("Last Usable LBA:\t\t%d\n", header.LastUsableLBA)
	fmt.Printf("Partition Entries Start LBA:\t%d\n", header.StartLBA)
	fmt.Printf("Partition Entry Size:\t\t%d\n", header.SizeOfSinglePartitionEntry)

	fmt.Println("==================")
}
