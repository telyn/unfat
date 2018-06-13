// dir implements support for FAT32 directory entries (including i haven
package dir

import (
	"encoding/binary"
	"fmt"
	"time"
)

const FileAttrRO = 0x01
const FileAttrHidden = 0x02
const FileAttrSystem = 0x04
const FileAttrVolumeLabel = 0x08
const FileAttrSubdirectory = 0x10
const FileAttrArchive = 0x20
const FileAttrDevice = 0x40

const fileAttrLFN = 0x0F

type File struct {
	ShortName    string
	LongName     string
	Attributes   uint8
	CreationTime time.Time
	AccessTime   time.Time
	ModifiedTime time.Time
	FirstCluster uint32
	Size         uint32

	shortNameChecksum uint16
	longNameBuf       []string
}

func lfnChecksum(str string) byte {
	sn := []byte(padShortName(str))

	sum := byte(0)

	for i := 11; i != 0; i-- {
		sum = ((sum & 1) << 7) +
			(sum >> 1) + sn[i]

	}
	return sum
}

// repeatedly call UnmarshalBinary with a 32-byte directory entry
func (f *File) UnmarshalBinary(buf []byte) (err error) {
	if len(buf) != 32 {
		return fmt.Errorf("32-byte slice must be given, but was of length %d", len(buf))
	}

	const oName = 0x00
	const oAttributes = 0x0B
	const oCreateTimeFine = 0x0D
	const oCreateTime = 0x0E
	const oCreateDate = 0x10
	const oAccessDate = 0x12
	const oFirstClusterHigh = 0x14
	const oModifiedTime = 0x16
	const oModifiedDate = 0x18
	const oFirstClusterLow = 0x1A
	const oSize = 0x1C

	const oLFNSequenceNumber = 0x00
	const oLFNChars1 = 0x01
	const oLFNShortNameChecksum = 0x0D
	const oLFNChars2 = 0x0E
	const oLFNChars3 = 0x1C

	f.Attributes = buf[oAttributes]
	f.FirstCluster = uint32(buf[oFirstClusterLow]) << 8
	f.Size = binary.LittleEndian.Uint32(buf[oSize : oSize+3])
	shortNameChecksum := binary.LittleEndian.Uint16(buf[oLFNShortNameChecksum : oLFNShortNameChecksum+1])

	if f.Attributes == fileAttrLFN && f.Size != 0 && f.FirstCluster == 0 {
		// almost certainly a LFN
		if f.shortNameChecksum != 0 && f.shortNameChecksum != shortNameChecksum {
			return fmt.Errorf("LFN checksum disagreed: expected %x, got %x", f.shortNameChecksum, shortNameChecksum)
		}

		isFirst := (buf[oLFNSequenceNumber] & 0x20) != 0
		//sequenceNum := buf[oLFNSequenceNumber] & 0xF
		if isFirst && f.LongName != "" {
			return fmt.Errorf("Found multiple first LFN entries")
		}

		//idx := sequenceNum * 13

	} else {

	}
	return
}