package main

type FSInformationSector struct {
	Signature                  [4]byte // offset 0x000 // 0x52 0x52 0x61 0x41 (RRaA)
	Signature2                 [4]byte // offset 0x1E4 // 0x72 0x72 0x41 0x61 (rrAa)
	FreeClusters               uint32  // offset 0x1E8 // 0xFFFFFFFF if unknown
	MostRecentAllocatedCluster uint32  // offset 0x1EC // 0xFFFFFFFF if unknown
	Signature3                 [4]byte // offset 0x01FC // 0x00 0x00 0x55 0xAA
}
