package mscfb

// ========================================================================= //

const (
	HeaderLen               int = 512 // length of CFB file header, in bytes
	DirEntryLen             int = 128 // length of directory entry, in bytes
	NumDifatEntriesInHeader int = 109
)

// MagicNumber of CFB file header values:
var MagicNumber = []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}

const (
	ByteOrderMark    uint16 = 0xfffe
	MiniSectorShift  uint16 = 6 // 64-byte mini sectors
	MiniSectorLen    int    = 1 << (MiniSectorShift)
	MiniStreamCutoff uint32 = 4096
)

// Constants for FAT entries:
const (
	MaxRegularSector uint32 = 0xfffffffa
	InvalidSector    uint32 = 0xfffffffb
	DifatSector      uint32 = 0xfffffffc
	FatSector        uint32 = 0xfffffffd
	EndOfChain       uint32 = 0xfffffffe
	FreeSector       uint32 = 0xffffffff
)

// Constants for directory entries:
const (
	RootDirName               = "Root Entry"
	ObjTypeUnallocated uint8  = 0
	ObjTypeStorage     uint8  = 1
	ObjTypeStream      uint8  = 2
	ObjTypeRoot        uint8  = 5
	ColorRed           uint8  = 0
	ColorBlack         uint8  = 1
	RootStreamId       uint32 = 0
	NoStream           uint32 = 0xffffffff
)
