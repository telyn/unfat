// Package dir implements support for FAT32 directory entries
package dir

import (
	"encoding/binary"
	"time"

	"github.com/telyn/unfat/fat/dir/attrs"
	"github.com/telyn/unfat/fat/dir/lfn"
)

const fileAttrLFN = 0x0F

type File struct {
	Name         string
	Attributes   attrs.Attributes
	CreationTime time.Time
	AccessTime   time.Time
	ModifiedTime time.Time
	FirstCluster uint32
	Size         uint32
	LFN          lfn.LongFileName
}

// ReadDirectoryEntry attempts to read a full directory entry starting at the
// beginning of buf, including a Long File Name if there is one.
// If an LFN with no matching directory entry is found, returns a File
// representing the long name and an error.
// numEntries is always guaranteed to be the number of entries used during the
// read, regardless of any error conditions. In other words - these two calls
// will always read separate files:
// f1, n, err := ReadDirectoryEntry(buf)
// f2, n, err := ReadDirectoryEntry(buf[32*n:])
func ReadDirectoryEntry(buf []byte) (f File, numEntries int, err error) {
	// first try to read an LFN
	lfn, numEntries, err := lfn.ReadLongFileName(buf)
	f.LFN = lfn
	if err != nil {
		return
	}

	buf = buf[numEntries*32:]

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

	f.Name = unpadShortName(buf[oName : oName+11])
	f.Attributes = attrs.ReadAttributes(buf[oAttributes], buf[oAttributes+1])
	// fuck about with time???
	f.FirstCluster = uint32(buf[oFirstClusterLow])<<8 | oFirstClusterHigh
	f.Size = binary.LittleEndian.Uint32(buf[oSize : oSize+3])
	return
}
