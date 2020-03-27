package wavreader

import (
	"fmt"
	"io"
)

func readU8(r io.ReaderAt, off int64) (uint8, error) {
	b := make([]byte, 1)
	n, err := r.ReadAt(b, off)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, fmt.Errorf("could not read byte at %v", err)
	}
	return b[0], nil
}

func readU16(r io.ReaderAt, off int64) (uint16, error) {
	b := make([]byte, 2)
	n, err := r.ReadAt(b, off)
	if err != nil {
		return 0, err
	}
	if n != 2 {
		return 0, fmt.Errorf("could not read byte at %v", err)
	}
	return uint16(b[0]) | uint16(b[1])<<8, nil
}

func readS16(r io.ReaderAt, off int64) (int16, error) {
	u16, err := readU16(r, off)
	if err != nil {
		return 0, err
	}
	return int16(u16), nil
}

func readU32(r io.ReaderAt, off int64) (uint32, error) {
	b := make([]byte, 4)
	n, err := r.ReadAt(b, off)
	if err != nil {
		return 0, err
	}
	if n != 4 {
		return 0, fmt.Errorf("could not read byte at %v", err)
	}
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24, nil
}

func readS32(r io.ReaderAt, off int64) (int32, error) {
	u32, err := readU32(r, off)
	if err != nil {
		return 0, err
	}
	return int32(u32), nil
}

func readU64(r io.ReaderAt, off int64) (uint64, error) {
	b := make([]byte, 8)
	n, err := r.ReadAt(b, off)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, fmt.Errorf("could not read byte at %v", err)
	}
	return uint64(b[0]) | uint64(b[1])<<8 |
		uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[2])<<32 | uint64(b[3])<<40 |
		uint64(b[2])<<48 | uint64(b[3])<<56, nil
}

func readS64(r io.ReaderAt, off int64) (int64, error) {
	u64, err := readU64(r, off)
	if err != nil {
		return 0, err
	}
	return int64(u64), nil
}
