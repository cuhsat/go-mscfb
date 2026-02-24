package mscfb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Header struct {
	Version            Version
	NumDirSectors      uint32
	NumFatSectors      uint32
	FirstDirSector     uint32
	FirstMinifatSector uint32
	NumMinifatSector   uint32
	FirstDifatSector   uint32
	NumDifatSectors    uint32

	InitialDifatEntries []uint32
}

const (
	reservedAfterMagicNumber = 16
	reservedAfterMiniShift   = 6
)

func (h *Header) readFrom(reader io.ReadSeeker) error {
	magicPart := make([]byte, len(MagicNumber))
	_, err := reader.Read(magicPart)
	if err != nil {
		return err
	}

	if !bytes.Equal(magicPart, MagicNumber) {
		return ErrorInvalidCFB
	}

	// seek reserved field
	_, err = reader.Seek(reservedAfterMagicNumber, io.SeekCurrent)
	if err != nil {
		return err
	}

	var minorVersion uint16
	err = binary.Read(reader, binary.LittleEndian, &minorVersion)
	if err != nil {
		return err
	}

	var versionNumber uint16
	err = binary.Read(reader, binary.LittleEndian, &versionNumber)
	if err != nil {
		return err
	}

	var byteOrderMark uint16
	err = binary.Read(reader, binary.LittleEndian, &byteOrderMark)
	if err != nil {
		return err
	}

	if byteOrderMark != ByteOrderMark {
		return fmt.Errorf("invalid CFB byte order mark (expected %x, found %x)", ByteOrderMark, byteOrderMark)
	}

	version, err := VersionNumber(versionNumber)
	if err != nil {
		return err
	}

	var sectorShift uint16
	err = binary.Read(reader, binary.LittleEndian, &sectorShift)
	if err != nil {
		return err
	}
	if sectorShift != version.SectorShift() {
		return fmt.Errorf("incorrect sector shift for CFB version %v (expected %v, found %v)", version, version.SectorShift(), sectorShift)
	}

	var miniSectorShift uint16
	err = binary.Read(reader, binary.LittleEndian, &miniSectorShift)
	if err != nil {
		return err
	}
	if miniSectorShift != MiniSectorShift {
		return fmt.Errorf("incorrect mini sector shift (expected %v, found %v)", MiniSectorShift, miniSectorShift)
	}

	// seek reserved field
	_, err = reader.Seek(reservedAfterMiniShift, io.SeekCurrent)
	if err != nil {
		return err
	}

	var numDirSectors uint32
	var numFatSectors uint32
	var firstDirSector uint32
	var transactionSign uint32

	err = binary.Read(reader, binary.LittleEndian, &numDirSectors)
	if err != nil {
		return err
	}

	err = binary.Read(reader, binary.LittleEndian, &numFatSectors)
	if err != nil {
		return err
	}

	err = binary.Read(reader, binary.LittleEndian, &firstDirSector)
	if err != nil {
		return err
	}

	err = binary.Read(reader, binary.LittleEndian, &transactionSign)
	if err != nil {
		return err
	}

	var miniStreamCutoff uint32
	err = binary.Read(reader, binary.LittleEndian, &miniStreamCutoff)
	if err != nil {
		return err
	}
	if miniStreamCutoff != MiniStreamCutoff {
		return fmt.Errorf("incorrect mini stream cutoff (expected %v, found %v)", MiniStreamCutoff, miniStreamCutoff)
	}

	var firstMinifatSector uint32
	var numMinifatSectors uint32
	var firstDifatSector uint32
	var numDifatSectors uint32

	err = binary.Read(reader, binary.LittleEndian, &firstMinifatSector)
	if err != nil {
		return err
	}

	err = binary.Read(reader, binary.LittleEndian, &numMinifatSectors)
	if err != nil {
		return err
	}

	err = binary.Read(reader, binary.LittleEndian, &firstDifatSector)
	if err != nil {
		return err
	}

	err = binary.Read(reader, binary.LittleEndian, &numDifatSectors)
	if err != nil {
		return err
	}

	// Some CFB implementations use FREE_SECTOR to indicate END_OF_CHAIN.
	if firstDifatSector == FreeSector {
		firstDifatSector = EndOfChain
	}

	difatEntries := make([]uint32, NumDifatEntriesInHeader)

	for i := range difatEntries {

		var next uint32
		err = binary.Read(reader, binary.LittleEndian, &next)
		if err != nil {
			return err
		}

		if next == FreeSector {
			break
		} else if next > MaxRegularSector {
			return fmt.Errorf("invalid DIFAT entry (expected value <= %v, found %v)", MaxRegularSector, next)

		}
		difatEntries[i] = next
	}

	h.Version = version
	h.NumDirSectors = numDirSectors
	h.NumFatSectors = numFatSectors
	h.FirstDirSector = firstDirSector
	h.FirstMinifatSector = firstMinifatSector
	h.NumMinifatSector = numMinifatSectors
	h.FirstDifatSector = firstDifatSector
	h.NumDifatSectors = numDifatSectors
	h.InitialDifatEntries = difatEntries

	return nil
}
