package dir

// Checksum returns the checksum for an 8.3 name
func Checksum(eightThree []byte) (sum byte) {
	for _, b := range eightThree {
		sum = ((sum & 1) << 7) + (sum >> 1) + b
	}

	return sum
}
