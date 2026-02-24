package mscfb

import "fmt"

const (
	V3 Version = 3
	V4 Version = 4
)

type Version int

func VersionNumber(v uint16) (Version, error) {
	switch v {
	case 3:
		return V3, nil
	case 4:
		return V4, nil
	default:
		return 0, fmt.Errorf("invalid version number: %v", v)
	}
}

func (v Version) SectorShift() uint16 {
	return uint16(v * 3)
}

func (v Version) SectorLen() int {
	return 1 << (int(v.SectorShift()))
}

func (v Version) SectorLenMask() uint64 {
	switch v {
	case V3:
		return 0xffffffff
	case V4:
		return 0xffffffffffffffff
	default:
		return 0
	}
}

func (v Version) DirEntriesPerSector() int {
	return v.SectorLen() / DirEntryLen
}
