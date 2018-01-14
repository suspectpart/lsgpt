package mbr

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"
)

const (
	_MBRSize = 512
)

// PartitionEntry represents one PartitonEntry in a MBR Partition Table
type PartitionEntry struct {
	Status          byte
	FirstSectorCHS  [3]byte
	PartitionType   byte
	LastSectorCHS   [3]byte
	FirstSectorLBA  uint32
	NumberOfSectors uint32
}

// MBR represents the Master Boot Record
type MBR struct {
	BootstrapCode    [446]byte
	PartitionEntries [4]PartitionEntry
	BootSignature    uint16
}

// ReadFrom reads MBR from file
func ReadFrom(filename string) (*MBR, error) {
	mbr := MBR{}

	file, err := os.Open(filename)

	if err != nil {
		return &mbr, err
	}

	defer file.Close()

	tableBytes := make([]byte, _MBRSize)

	file.Read(tableBytes)

	tableBuffer := bytes.NewBuffer(tableBytes)

	_ = binary.Read(tableBuffer, binary.LittleEndian, &mbr)

	if mbr.BootSignature != 0xaa55 {
		return &MBR{}, errors.New("No MBR found on " + filename)
	}

	return &mbr, nil
}

// IsProtective returns whether the MBR is protective (to secure a GPT on older BIOS systems)
func (mbr *MBR) IsProtective() bool {
	return mbr.PartitionEntries[0].PartitionType == 0xee
}
