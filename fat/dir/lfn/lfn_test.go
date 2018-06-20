package dir

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

type tw struct {
	t *testing.T
}

func (tw tw) Write(bytes []byte) error {
	tw.t.Log(string(bytes))
	return nil
}

func TestReadLongFileName(t *testing.T) {
	stdout := os.Stdout
	bytes := helperLoadDirectoryEntries(t)
	tests := []struct {
		name             string
		offset           int
		numEntries       int
		shouldErr        bool
		errMessageRegexp *regexp.Regexp
		checksum         byte
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
	}, {
		name:             "error-bad-sequence",
		numEntries:       2,
		offset:           0x1E0,
		shouldErr:        true,
		errMessageRegexp: regexp.MustCompile("(?i:sequence number mismatch)"),
	}, {
		name:             "error-bad-checksum-",
		numEntries:       3,
		offset:           0x140,
		shouldErr:        true,
		errMessageRegexp: regexp.MustCompile("(?i:checksum mismatch)"),
	}, {
		name:       "error-interrupted",
		numEntries: 1,
		offset:     0x1a0,
		shouldErr:  true,
	}, {
		name:       "error-empty-file",
		numEntries: 1,
		offset:     0x1e0,
		shouldErr:  true,
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w
			lfn, numEntries, err := ReadLongFileName(bytes[test.offset:])
			w.Close()
			b, _ := ioutil.ReadAll(r)
			t.Log(string(b))

			if err != nil && !test.shouldErr {
				t.Fatal(err)
			} else if err == nil && test.shouldErr {
				t.Fatal("Should have errored but didn't")
			}
			if test.shouldErr && err != nil {
				if test.errMessageRegexp != nil &&
					!test.errMessageRegexp.MatchString(err.Error()) {
					t.Errorf("err message didn't match %v. Got %q", test.errMessageRegexp, err.Error())
				}
				return
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
	os.Stdout = stdout
}
