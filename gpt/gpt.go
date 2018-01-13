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

// ReadHeader reads the GPT header from a file
func ReadHeader(filename string) (*Header, error) {
	header := Header{}

	file, err := os.Open(filename)

	if err != nil {
		return &header, err
	}

	defer file.Close()

	headerBytes := make([]byte, 512)

	file.ReadAt(headerBytes, _GPTHeaderOffset)

	headerBuffer := bytes.NewBuffer(headerBytes)

	_ = binary.Read(headerBuffer, binary.LittleEndian, &header)

	if string(header.Signature[:8]) != _GPTSignature {
		return &Header{}, errors.New("No GPT found on " + filename)
	}

	return &header, nil
}

func (header *Header) String() string {
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
	)
}
