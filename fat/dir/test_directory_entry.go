package dir_test

import (
	"github.com/telyn/unfat/fat/dir"
	"io/ioutil"
	"testing"
)

func helperLoadDirectoryEntries() ([]byte, error) {
	return ioutil.ReadAll("testdata/directory-entries.fat32")
}

func TestDirectoryEntry(t *testing.T) {
	data, err := helperLoadDirectoryEntries()
	tests := []struct {
		// start of entry offset in dataset
		entryStart int
		// entries to pump into the file
		entries   uint
		filename  string
		shortname string
		firstCluser uint32
		size uint32
	}{{
		name: "simple shortname entry",
		startOffset: 0x20,
		entries: 1
		filename: "TEST-DIR",
		shortname: "TEST-DIR",
		startCluster: 0x34CB1,
		size: 0,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := dir.File{}
			for offset := 0; offset < 32*test.entries; i+= 32 {
				data := data[offset:offset+32]
				f.UnmarshalBinary(data)
			}
			if f.ShortName != test.shortname {
				t.Errorf("ShortName expected to be %q but was %q", test.shortname, f.ShortName)
			}
		})
	}
}
