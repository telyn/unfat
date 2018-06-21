package dir

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

func ReadAttributes(b1 byte, b2) (a Attributes) {
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
}
