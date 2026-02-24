package mscfb

type ObjectType int

const (
	ObjUnallocated ObjectType = iota
	ObjStorage
	ObjStream
	ObjRoot
)

func (o ObjectType) AsByte() byte {
	switch o {
	case ObjUnallocated:
		return ObjTypeUnallocated
	case ObjStorage:
		return ObjTypeStorage
	case ObjStream:
		return ObjTypeStream
	case ObjRoot:
		return ObjTypeRoot
	default:
		return 0
	}
}

func ObjectFromByte(b byte) ObjectType {
	switch b {
	case ObjTypeUnallocated:
		return ObjUnallocated
	case ObjTypeStorage:
		return ObjStorage
	case ObjTypeStream:
		return ObjStream
	case ObjTypeRoot:
		return ObjRoot
	default:
		return -1
	}
}
