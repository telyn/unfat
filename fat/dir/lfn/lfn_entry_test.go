package lfn

import (
	"reflect"
	"testing"
)

func TestReadLFNEntry(t *testing.T) {
	tests := []struct {
		name     string
		offset   int
		expected LFNEntry
	}{{
		name:   "test-dir",
		offset: 0,
		expected: LFNEntry{
			SequenceNumber: 1,
			IsLast:         true,
			Chars: []byte{
				't', 0, 'e', 0, 's', 0, 't', 0,
				'-', 0, 'd', 0, 'i', 0, 'r', 0,
				0, 0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF,
			},
		},
	}, {
		name:   "test-dir2",
		offset: 0x40,
		expected: LFNEntry{
			SequenceNumber: 1,
			IsLast:         true,
			Chars: []byte{
				't', 0, 'e', 0, 's', 0, 't', 0,
				'-', 0, 'd', 0, 'i', 0, 'r', 0,
				'2', 0, 0, 0, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF,
			},
		},
	}, {
		name:   "big-file",
		offset: 0x80,
		expected: LFNEntry{
			SequenceNumber: 1,
			IsLast:         true,
			Chars: []byte{
				'b', 0, 'i', 0, 'g', 0, '-', 0,
				'f', 0, 'i', 0, 'l', 0, 'e', 0,
				0, 0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF,
			},
		},
	}, {
		name:   "gifjpegtxt",
		offset: 0xc0,
		expected: LFNEntry{
			SequenceNumber: 3,
			IsLast:         true,
			Chars: []byte{
				'g', 0, 'i', 0, 'f', 0, 'j', 0,
				'p', 0, 'e', 0, 'g', 0, 't', 0,
				'x', 0, 't', 0, 0, 0, 0xFF, 0xFF,
				0xFF, 0xFF,
			},
		},
	}, {
		name:   "-file-waddup.",
		offset: 0xe0,
		expected: LFNEntry{
			SequenceNumber: 2,
			IsLast:         false,
			Chars: []byte{
				'-', 0, 'f', 0, 'i', 0, 'l', 0,
				'e', 0, '-', 0, 'w', 0, 'a', 0,
				'd', 0, 'd', 0, 'u', 0, 'p', 0,
				'.', 0,
			},
		},
	}, {
		name:   "yeah-file-big",
		offset: 0x100,
		expected: LFNEntry{
			SequenceNumber: 1,
			IsLast:         false,
			Chars: []byte{
				'y', 0, 'e', 0, 'a', 0, 'h', 0,
				'-', 0, 'f', 0, 'i', 0, 'l', 0,
				'e', 0, '-', 0, 'b', 0, 'i', 0,
				'g', 0,
			},
		},
	}}
	for _, test := range tests {
		bytes := helperLoadDirectoryEntries(t)
		t.Run(test.name, func(t *testing.T) {
			entry, err := readLFNEntry(bytes[test.offset:])
			if err != nil {
				t.Fatal(err)
			}
			if entry.SequenceNumber != test.expected.SequenceNumber {
				t.Errorf("Sequence number was wrong. Expected %d, got %d", test.expected.SequenceNumber, entry.SequenceNumber)
			}
			if entry.IsLast != test.expected.IsLast {
				t.Errorf("IsLast was wrong, expected %v, got %v", test.expected.IsLast, entry.IsLast)
			}
			if !reflect.DeepEqual(test.expected.Chars, entry.Chars) {
				t.Errorf("Chars was wrong, expected %x, got %x", test.expected.Chars, entry.Chars)
			}
		})
	}
}
