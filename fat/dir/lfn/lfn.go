package lfn

import (
	"encoding/binary"
	"fmt"
	"unicode/utf16"
)

type LongFileName struct {
	Name     string
	Checksum byte
}

func (lfn LongFileName) Matches(eightThree []byte) bool {
	return Checksum(eightThree) == lfn.Checksum
}

// ReadLongFileName reads a full VFAT LFN from the directory in bytes.
// At the end of reading, lfn is the filename that was managed to be made,
// numEntries is the number of entries that were read, and err is nil for
// success or something else otherwise.
// numEntries*32 is guaranteed to always be the offset of the next unread entry
// in the directory. Some examples:
//
//   * if bytes[0:32] is not an LFN, numEntries = 0. Try reading bytes[0:32] as
//     a normal file entry
//   * if bytes[0:32] does appear to be an LFN entry but is not the end of a
//     sequence, numEntries = 1
//   * if bytes[0:32] is the last entry in an LFN sequence but bytes[32:64] has
//     an incorrect sequence number, numEntries = 1
//   * if bytes[0:32] is the last entry in an LFN sequence and so is
//     bytes[32:64], numEntries = 1 so that a valid LFN can be attempted from
//     bytes[32:]
//
// if err is not nil, lfn is to be assumed to be invalid. It may contain some
// info about the file name or may not. It will always only be partial at best
// though.
func ReadLongFileName(bytes []byte) (lfn LongFileName, numEntries int, err error) {
	entry, err := readLFNEntry(bytes)
	if err != nil {
		if _, ok := err.(NotLFNEntry); ok {
			err = nil
		}
		return
	}
	numEntries++
	if !entry.IsLast {
		err = fmt.Errorf("First LFN entry read was not final in its sequence")
		return
	}
	fmt.Println("last entry ok")

	entries := make([]LFNEntry, entry.SequenceNumber)
	entries[entry.SequenceNumber-1] = entry
	lfn.Checksum = entry.Checksum
	for i := entry.SequenceNumber - 1; i > 0; i-- {
		if len(bytes) < 32 {
			err = fmt.Errorf("Unexpected end of directory")
			return
		}
		bytes = bytes[32:]

		entry, err = readLFNEntry(bytes)
		if err != nil {
			return
		}
		if entry.IsLast {
			err = fmt.Errorf("LFN entry never finishes before next entry begins")
			return
		}
		fmt.Printf("%d entry ok\n", i)
		numEntries++
		if entry.SequenceNumber != i {
			err = fmt.Errorf("sequence number mismatch - expected %d, got %d",
				i, entry.SequenceNumber)
			return
		}

		if entry.Checksum != lfn.Checksum {
			err = fmt.Errorf("checksum mismatch for sequence number %d: %x "+
				"(expecting %x)", i, entry.Checksum, lfn.Checksum)
			return
		}
		entries[i-1] = entry
	}
	lfn.Name = parseEntries(entries)
	return
}

func parseEntries(entries []LFNEntry) (name string) {
	ints := make([]uint16, 0, 13*len(entries))
	for _, entry := range entries {
		for c := 0; c < 26; c += 2 {
			char := binary.LittleEndian.Uint16(entry.Chars[c:])
			if char == 0 {
				break
			}
			ints = append(ints, char)
		}
	}

	runes := utf16.Decode(ints)
	return string(runes)
}
