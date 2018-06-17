package dir

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
	"unicode/utf16"
)

func helperLoadDirectoryEntries(t *testing.T) []byte {
	bytes, err := ioutil.ReadFile("./testdata/directory-entries.fat32")
	if err != nil {
		t.Fatalf("Couldn't load test data: %v", err)
	}
	return bytes
}

func ceil(nom, denom int) int {
	base := nom / denom
	if nom%denom > 0 {
		return base + 1
	}
	return base
}

func pad(bit []byte) (out []byte) {
	out = bit
	if len(out) < 26 {
		for i := len(out); i < 26; i++ {
			out = append(out, 0xFF)
		}
	}
	return
}

func ucs2sgl(name string) (bytes []byte) {
	if len(name) > 13 {
		return
	}
	bytes = make([]byte, len(name)*2, 26)
	ints := utf16.Encode([]rune(name))
	fmt.Printf("ints: %#v", ints)
	for i := range ints {
		binary.LittleEndian.PutUint16(bytes[i*2:], ints[i])
	}
	if len(bytes) < 26 {
		bytes = append(bytes, 0x00, 0x00)
	}
	if len(bytes) < 26 {
		for i := len(bytes); i < 26; i++ {
			bytes = append(bytes, 0xFF)
		}
	}
	return
}

func TestUcs2sgl(t *testing.T) {
	tests := []struct {
		in       string
		expected []byte
	}{{
		in: "hello",
		expected: []byte{'h', 0, 'e', 0, 'l', 0, 'l', 0, 'o', 0,
			0, 0,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	}}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			res := ucs2sgl(test.in)
			if !reflect.DeepEqual(test.expected, res) {
				t.Errorf("expect: %x\nactual %x", test.expected, res)
			}
		})
	}
}

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
