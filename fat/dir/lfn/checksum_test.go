package dir_test

import (
	"testing"

	dir "github.com/telyn/unfat/fat/dir/lfn"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		in       string
		expected byte
	}{{
		in:       "TEST-DIR   ",
		expected: 0x78,
	}, {
		in:       "TEST-D~1   ",
		expected: 0xC7,
	}, {
		in:       "BIG-FILE   ",
		expected: 0xD7,
	}, {
		in:       "YEAH-F~1GIF",
		expected: 0xE2,
	}}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			if dir.Checksum([]byte(test.in)) != test.expected {
				t.Errorf("Checksum(%q) != %x", test.in, test.expected)
			}
		})
	}
}
