package dir_test

import (
	"io/ioutil"
	"reflect"
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

func assertEqual(t *testing.T, fmt string, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(fmt, expected, actual)
	}
}

func TestDirectoryEntry(t *testing.T) {
	data := helperLoadDirectoryEntries(t)
	tests := []struct {
		name string
		// offset of first entry for this file
		offset int
		// entries to pump into the file
		numEntries int
		expected   dir.File
		shouldErr  bool
	}{{
		name:       "simple shortname entry",
		offset:     0x20,
		numEntries: 1,
		expected: dir.File{
			Name:         "TEST-DIR",
			FirstCluster: 0x34CB1,
			Size:         0,
			// TODO: times :-(
		},
	}, {
		name:       "simple lfn entry",
		offset:     0x00,
		numEntries: 2,
		expected:   dir.File{},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, numEntries, err := dir.ReadDirectoryEntry(data[test.offset:])
			if !test.shouldErr && err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			} else if test.shouldErr && err == nil {
				t.Errorf("Expected error but did not get one")
			}
			assertEqual(t, "expected numEntries == %d, was %d", test.numEntries, numEntries)
			if !test.shouldErr {
				assertEqual(t, "expected Name == %q, was %q", test.expected.Name, file.Name)
				assertEqual(t, "expected Attributes == %v, was %v", test.expected.Attributes, file.Attributes)
				assertEqual(t, "expected CreationTime == %v, was %v", test.expected.CreationTime, file.CreationTime)
				assertEqual(t, "expected AccessTime == %v, was %v", test.expected.AccessTime, file.AccessTime)
				assertEqual(t, "expected ModifiedTime == %v, was %v", test.expected.ModifiedTime, file.ModifiedTime)
				assertEqual(t, "expected FirstCluster == %v, was %v", test.expected.FirstCluster, file.FirstCluster)
				assertEqual(t, "expected Size == %v, was %v", test.expected.Size, file.Size)
				assertEqual(t, "expected LFN.Name == %v, was %v", test.expected.LFN.Name, file.LFN.Name)
				assertEqual(t, "expected LFN.Checksum == %v, was %v", test.expected.LFN.Checksum, file.LFN.Checksum)
			}
		})
	}
}
