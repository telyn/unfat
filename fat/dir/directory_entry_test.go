package dir_test

import (
	"io/ioutil"
	"testing"

	"github.com/telyn/unfat/fat/dir"
)

func helperLoadDirectoryEntries(t *testing.T) []byte {
	bytes, err := ioutil.ReadFile("./testdata/directory-entries.fat32")
	if err != nil {
		t.Fatalf("Couldn't load test data: %v", err)
	}
	return bytes
}

func TestDirectoryEntry(t *testing.T) {
	helperLoadDirectoryEntries(t)
	tests := []struct {
		name string
		// start of entry offset in dataset
		startOffset int
		// entries to pump into the file
		entries      uint
		filename     string
		shortname    string
		firstCluster uint32
		size         uint32
	}{{
		name:         "simple shortname entry",
		startOffset:  0x20,
		entries:      1,
		filename:     "TEST-DIR",
		shortname:    "TEST-DIR",
		firstCluster: 0x34CB1,
		size:         0,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := dir.File{}
			for offset := uint(0); offset < 32*test.entries; offset += 32 {
				//data := data[offset : offset+32]
				//f.UnmarshalBinary(data)
			}
			if f.Name != test.shortname {
				//t.Errorf("ShortName expected to be %q but was %q", test.shortname, f.ShortName)
			}
		})
	}
}
