package mscfb

type Color int

const (
	Red Color = iota
	Black
)

func (c Color) AsByte() byte {
	switch c {
	case Red:
		return ColorRed
	case Black:
		return ColorBlack
	default:
		return 0
	}
}

func ColorFromByte(b byte) Color {
	switch b {
	case ColorRed:
		return Red
	case ColorBlack:
		return Black
	default:
		return -1
	}
}
