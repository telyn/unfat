package main

type BootSector struct {
	Jump    [3]byte // offset 0x000
	OEMName [8]byte // offset 0x003
	// DOS 2.0 BIOS Parameter Block
	BytesPerLogicalSector    uint16 // offset 0x00B // number that is a power of 2
	LogicalSectorsPerCluster uint8  // offset 0x00D // 1, 2, 4, 8, 16, 32, 64, and 128
	ReservedSectors          uint16 // offset 0x00E //
	NumFATs                  uint8  // offset 0x010 // usually 2
	MaxRootDirEntries        uint16 // offset 0x011 // 0 for FAT32
	TotalLogicalSectors      uint8  // offset 0x013 (if 0, use LongTotalLogicalSectors)
	MediaDescriptor          uint8  // offset 0x015 // probably F8 (Fixed Disk) // prob ignore it
	LogicalSectorsPerFAT     uint16 // offset 0x016 // if 0, use LongLogicalSectorsPerFat
	// DOS 3.31 BIOS Parameter Block
	PhysicalSectorsPerTrack      uint16 // offset 0x018 // if 0 - reserved but not used // 0 or 1 prob means LBA // prob ignore it
	NumHeads                     uint16 // offset 0x01A // if 0 - reserved but not used // 0 or 1 prob means LBA // prob ignore it
	HiddenSectorsBeforePartition uint16 // offset 0x01C // prob ignore it
	LongTotalLogicalSectors      uint32 // offset 0x020
	// FAT12/FAT16 Extended BPB
	ExtendedBootSignature uint8    // offset 0x026 // 0x29 if next 3 values exist, 0x28 if only VolumeID does // FAT32 one is at 0x042
	VolumeID              uint32   // offset 0x027 // FAT32 @ 0x043
	VolumeLabel           [11]byte // offset 0x02B // FAT32 @ 0x047
	FSType                [8]byte  // offset 0x036 // FAT / FAT12 / FAT16 // FAT32 @ 0x052
	// FAT32 Extended BPB
	LongLogicalSectorsPerFat  uint32 // offset 0x024 // avoid any number with 0x28 or 0x29 in the byte @ 0x026
	Version                   uint16 // offset 0x02A // high byte 0x2B
	ClusterNumberOfRootDir    uint32 // offset 0x02C
	FSInformationSector       uint16 // offset 0x030 // 0x0000 or 0xFFFF means don't go looking
	BootSectorCopyFirstSector uint16 // offset 0x032 // prob ignore it

	Signature [2]byte // offset 0x1FE // 0x55 0xAA for FAT
}
