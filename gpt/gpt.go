package gpt

import (
	"bytes"
	"encoding/binary"

	"github.com/google/uuid"
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
