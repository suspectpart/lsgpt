package gpt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

const (
	_GPTHeaderOffset = 512
	_GPTSignature    = "EFI PART"
)

// PartitionEntry represents one entry in the GPT Partition Array
type PartitionEntry struct {
	PartitionType     GUID
	UniquePartitionID GUID
	FirstLBA          uint64
	LastLBA           uint64
	Flags             uint64
	PartitonName      [72]byte
}

// GUID represents a GUID in binary format according to RFC4122 as in GPT Header
type GUID struct {
	TimeLow          uint32
	TimeMid          uint16
	TimeHiAndVersion uint16
	Nodes            uint64
}

// AsUUID takes a GUID structure and transforms it to a uuid.UUID
func (guid *GUID) AsUUID() uuid.UUID {
	var result uuid.UUID

	buf := new(bytes.Buffer)

	_ = binary.Write(buf, binary.BigEndian, guid.TimeLow)
	_ = binary.Write(buf, binary.BigEndian, guid.TimeMid)
	_ = binary.Write(buf, binary.BigEndian, guid.TimeHiAndVersion)
	_ = binary.Write(buf, binary.LittleEndian, guid.Nodes)

	copy(result[:], buf.Bytes())

	return result
}

// Header represents the header of a GPT at LBA1
type Header struct {
	Signature                  [8]byte
	Revision                   uint32
	HeaderSize                 uint32
	HeaderCRC32                uint32
	_                          uint32 // reserved
	CurrentLBA                 uint64
	BackupLBA                  uint64
	FirstUsableLBA             uint64
	LastUsableLBA              uint64
	DiskGUID                   GUID
	StartLBA                   uint64
	NumberOfPartitionEntries   uint32
	SizeOfSinglePartitionEntry uint32
	PartitionArrayCRC32        uint32
	_                          [420]byte
}

// GUIDPartitionTable represents a GPT on EFI systems
type GUIDPartitionTable struct {
	Header  Header
	Entries [128]PartitionEntry
}

// ReadFrom reads the GPT header from a file
func ReadFrom(filename string) (*GUIDPartitionTable, error) {
	table := GUIDPartitionTable{}

	file, err := os.Open(filename)

	if err != nil {
		return &table, err
	}

	defer file.Close()

	tableBytes := make([]byte, 512+128*128)

	file.ReadAt(tableBytes, _GPTHeaderOffset)

	tableBuffer := bytes.NewBuffer(tableBytes)

	_ = binary.Read(tableBuffer, binary.LittleEndian, &table)

	if string(table.Header.Signature[:8]) != _GPTSignature {
		return &GUIDPartitionTable{}, errors.New("No GPT found on " + filename)
	}

	return &table, nil
}

// IsEmpty checks if a partition entry does not point to an existing partition
func (entry PartitionEntry) IsEmpty() bool {
	return entry.FirstLBA == 0 && entry.LastLBA == 0
}

// NumPartitions returns the number of used partitions in the given GUIDPartitionTable
func (table *GUIDPartitionTable) NumPartitions() int {
	partitions := 0

	for _, entry := range table.Entries {
		if entry.IsEmpty() {
			break
		}

		partitions++
	}

	return partitions
}

func (table *GUIDPartitionTable) String() string {
	str := table.Header.String()

	str += fmt.Sprintf("Number\tStart (sector)\tEnd (sector)\tName\n")

	for i := 0; i < table.NumPartitions(); i++ {
		var entry = table.Entries[i]

		str += fmt.Sprintf("%d\t%d\t\t%d\t\t%s\n", i+1, entry.FirstLBA, entry.LastLBA, entry.PartitonName)
	}

	return str
}

func (header Header) String() string {
	return fmt.Sprintf(
		`=== GPT Header ===
Disk UUID:			%s
Header Checksum CRC32:		%d
Header Size:			%d
First Usable LBA:		%d
Last Usable LBA:		%d
Partition Entries Start LBA:	%d
Partition Entry Size:		%d
==================`,
		header.DiskGUID.AsUUID(),
		header.HeaderCRC32,
		header.HeaderSize,
		header.FirstUsableLBA,
		header.LastUsableLBA,
		header.StartLBA,
		header.SizeOfSinglePartitionEntry,
	) + "\n"
}
