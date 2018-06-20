package dir

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

func ReadLongFileName(bytes []byte) (lfn LongFileName, numEntries int, err error) {
	entry, err := readLFNEntry(bytes)
	if err != nil {
		return
	}
	numEntries++
	if !entry.IsLast {
		err = fmt.Errorf("First LFN entry read was not final in its sequence. Bailing")
		return
	}
	lfn.Checksum = entry.Checksum
	entries := make([]LFNEntry, entry.SequenceNumber)
	entries[entry.SequenceNumber-1] = entry
	for i := entry.SequenceNumber - 1; i > 0; i-- {
		bytes = bytes[32:]
		entry, err = readLFNEntry(bytes)
		if err != nil {
			return
		}
		if entry.IsLast {
			err = fmt.Errorf("LFN entry never finishes before next entry begins")
			return
		}
		numEntries++
		if entry.Checksum != lfn.Checksum {
			err = fmt.Errorf("Mismatched LFN entry checksum for sequence numbrer %d: %x (expecting %x)", i, entry.Checksum, lfn.Checksum)
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
