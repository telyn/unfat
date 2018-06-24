package lfn

import (
	"encoding/binary"
)

type LFNEntry struct {
	SequenceNumber int
	IsLast         bool
	Checksum       byte
	Chars          []byte
}

func readLFNEntry(bytes []byte) (entry LFNEntry, err error) {
	const oFirstChars = 0x01
	const oAttributes = 0x0B
	const oType = 0x0C
	const oChecksum = 0x0D
	const oSecondChars = 0x0E
	const oFirstCluster = 0x1A
	const oThirdChars = 0x1C
	const oEnd = 0x20

	if bytes[oAttributes] != 0xF {
		return entry, NotLFNEntry{Message: "attributes were not 0xF",
			Attributes: bytes[oAttributes]}
	}
	if bytes[oType] != 0x00 {
		return entry, NotLFNEntry{Message: "type was not 0x0",
			Attributes: bytes[oAttributes], Type: bytes[oType]}
	}
	firstCluster := binary.LittleEndian.Uint16(
		bytes[oFirstCluster : oFirstCluster+2])

	if firstCluster != 0 {
		return entry, NotLFNEntry{Message: "first cluster was not 0",
			Attributes: bytes[oAttributes], Type: bytes[oType],
			FirstCluster: firstCluster}
	}
	checksum := bytes[oChecksum]

	isLast := false
	if (bytes[0] & 0x40) == 0x40 {
		isLast = true
	}

	chars := make([]byte, 26)

	copy(chars[0:10], bytes[oFirstChars:oAttributes])
	copy(chars[10:22], bytes[oSecondChars:oFirstCluster])
	copy(chars[22:26], bytes[oThirdChars:oEnd])

	return LFNEntry{
		SequenceNumber: int(bytes[0] & 0xF),
		Checksum:       checksum,
		IsLast:         isLast,
		Chars:          chars}, nil
}
