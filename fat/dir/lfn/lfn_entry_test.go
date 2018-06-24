package lfn

import "testing"

func TestReadLFNEntry(t *testing.T) {
	tests := []struct {
		name     string
		offset   int
		expected LFNEntry
	}{{
		name:   "test-dir",
		offset: 0,
		expected: LFNEntry{
			SequenceNumber: 0,
			IsLast:         true,
			Chars: []byte{
				't', 0, 'e', 0, 's', 0, 't', 0,
				'-', 0, 'd', 0, 'i', 0, 'r', 0,
				0, 0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF},
		},
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})

	}
}
