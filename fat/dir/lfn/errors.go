package dir

import "fmt"

type NotLFNEntry struct {
	Attributes   byte
	Type         byte
	FirstCluster uint16
	Message      string
}

func (err NotLFNEntry) Error() string {
	return fmt.Sprintf("%s\nAttributes: %x, Type: %x, FirstCluster: %d",
		err.Message, err.Attributes, err.Type, err.FirstCluster)
}
