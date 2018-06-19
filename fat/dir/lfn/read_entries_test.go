package dir

import (
	"testing"
)

func TestReadLongFileName(t *testing.T) {
	bytes := helperLoadDirectoryEntries(t)
	tests := []struct {
		name       string
		offset     int
		numEntries int
		checksum   byte
	}{{
		name:       "test-dir",
		offset:     0,
		numEntries: 1,
		checksum:   0x78,
	}, {
		name:       "test-dir2",
		offset:     0x40,
		numEntries: 1,
		checksum:   0xC7,
	}, {
		name:       "big-file",
		offset:     0x80,
		numEntries: 1,
		checksum:   0xD7,
	}, {
		name:       "yeah-file-big-file-waddup.gifjpegtxt",
		offset:     0xC0,
		numEntries: 3,
		checksum:   0xE2,
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lfn, numEntries, err := ReadLongFileName(bytes[test.offset:])
			if err != nil {
				t.Fatal(err)
			}
			if lfn.Name != test.name {
				t.Errorf("%q != %q", test.name, lfn.Name)
			}
			if lfn.Checksum != test.checksum {
				t.Errorf("Checksums don't match. Expected %x, got %x", test.checksum, lfn.Checksum)
			}
			if numEntries != test.numEntries {
				t.Errorf("number of read entries incorrect. %d should be %d", test.numEntries, numEntries)
			}
		})
	}
}
