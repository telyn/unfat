package attrs

import "strings"

const AttrRO = 0x01
const AttrHidden = 0x02
const AttrSystem = 0x04
const AttrVolumeLabel = 0x08
const AttrSubdirectory = 0x10
const AttrArchive = 0x20
const AttrDevice = 0x40
const AttrReserved = 0x80

type Attributes struct {
	ReadOnly     bool
	Hidden       bool
	System       bool
	VolumeLabel  bool
	Subdirectory bool
	Archive      bool
	Device       bool
	Reserved     bool
}

func ReadAttributes(b1, b2 byte) (a Attributes) {
	if b1&AttrRO == AttrRO {
		a.ReadOnly = true
	}
	if b1&AttrHidden == AttrHidden {
		a.Hidden = true
	}
	if b1&AttrSystem == AttrSystem {
		a.System = true
	}
	if b1&AttrVolumeLabel == AttrVolumeLabel {
		a.VolumeLabel = true
	}
	if b1&AttrSubdirectory == AttrSubdirectory {
		a.Subdirectory = true
	}
	if b1&AttrArchive == AttrArchive {
		a.Archive = true
	}
	if b1&AttrDevice == AttrDevice {
		a.Device = true
	}
	if b1&AttrReserved == AttrReserved {
		a.Reserved = true
	}
	return
}

func (a Attributes) String() string {
	attrs := make([]string, 0, 160)
	if a.ReadOnly {
		attrs = append(attrs, "ReadOnly")
	}
	if a.Hidden {
		attrs = append(attrs, "Hidden")
	}
	if a.System {
		attrs = append(attrs, "System")
	}
	if a.VolumeLabel {
		attrs = append(attrs, "VolumeLabel")
	}
	if a.Subdirectory {
		attrs = append(attrs, "Subdirectory")
	}
	if a.Archive {
		attrs = append(attrs, "Archive")
	}
	if a.Device {
		attrs = append(attrs, "Device")
	}
	if a.Reserved {
		attrs = append(attrs, "Reserved")
	}
	return strings.Join(attrs, ", ")
}
