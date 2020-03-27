package wavreader

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

// Reader ...
type Reader struct {
	r             io.ReaderAt
	bits          uint16
	format, chans uint16
	rate          uint32
	data, size    uint64
	samples       uint64
}

// New ...
func New(r io.ReaderAt) (*Reader, error) {
	// Is RIFF ?
	mark := make([]byte, 4)
	n, err := r.ReadAt(mark, 0)
	if err != nil {
		return nil, fmt.Errorf("could not read header: %v", err)
	}
	if n != 4 || !bytes.Equal(mark, []byte{'R', 'I', 'F', 'F'}) {
		return nil, fmt.Errorf("not a RIFF file")
	}

	// get file size
	fsize, err := readU32(r, 4)
	if err != nil {
		return nil, fmt.Errorf("could not get filesize: %v", err)
	}

	// Is WAVE
	head := make([]byte, 4)
	n, err = r.ReadAt(head, 8)
	if err != nil {
		return nil, fmt.Errorf("could not read header: %v", err)
	}
	if n != 4 || !bytes.Equal(head, []byte{'W', 'A', 'V', 'E'}) {
		return nil, fmt.Errorf("not a WAVE file")
	}

	// Has fmt chunk ?
	wfmt := make([]byte, 4)
	n, err = r.ReadAt(wfmt, 12)
	if err != nil {
		return nil, fmt.Errorf("could not read header: %v", err)
	}
	if n != 4 || !bytes.Equal(wfmt, []byte{'f', 'm', 't', ' '}) {
		return nil, fmt.Errorf("file has no fmt chunk")
	}

	hsize, err := readU32(r, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to get fmt size: %v", err)
	}
	if hsize < 8 {
		return nil, fmt.Errorf("unexpected chunk size %d: %v",
			hsize, err)
	}

	format, err := readU16(r, 20)
	if err != nil {
		return nil, fmt.Errorf("failed to get format")
	}
	switch format {
	case 1:
	// case 0xfffe: // s32le, f32le, f64le
	// case 6: // alaw
	// case 7: // mulaw
	default:
		return nil, fmt.Errorf("unknown file format %d", format)
	}

	chans, err := readU16(r, 22)
	if err != nil {
		return nil, fmt.Errorf("failed to get chans")
	}

	sampleRate, err := readU32(r, 24)
	if err != nil {
		return nil, fmt.Errorf("failed to get sampleRate")
	}

	bits, err := readU16(r, 34)
	if err != nil {
		return nil, fmt.Errorf("failed to get bits")
	}
	switch bits {
	case 8, 16:
	default:
		return nil, fmt.Errorf("unsupported sample depth")
	}

	fs64 := uint64(fsize)
	data := uint64(0)

	next := uint64(hsize + 20)
	for next < fs64 {
		chunk := make([]byte, 4)
		n, err = r.ReadAt(chunk, int64(next))
		if err != nil {
			return nil, fmt.Errorf("could not read chunk at %d: %v",
				next, err)
		}

		size, err := readU32(r, int64(next+4))
		if err != nil {
			return nil, fmt.Errorf("could not read chunk size: %v",
				err)
		}

		// fmt.Printf("%c: %d\n", chunk, size)
		if bytes.Equal(chunk, []byte{'d', 'a', 't', 'a'}) {
			data = next + 8
			fs64 = uint64(size)
		}
		next += 8 + uint64(size)
	}

	if data == 0 {
		return nil, fmt.Errorf("file has no data")
	}

	count := fs64 / (uint64(bits) * uint64(chans) / 8)
	return &Reader{
		r:       r,
		format:  format,
		chans:   chans,
		data:    data,
		size:    fs64,
		bits:    bits,
		rate:    sampleRate,
		samples: count,
	}, nil
}

// Len ...
func (r *Reader) Len() uint64 {
	return r.samples
}

// Rate ...
func (r *Reader) Rate() uint32 {
	return r.rate
}

// Chans ...
func (r *Reader) Chans() uint16 {
	return r.chans
}

// Duration ...
func (r *Reader) Duration() time.Duration {
	s := float64(r.samples) / float64(r.rate)
	return time.Duration(s * float64(time.Second))
}

// At ...
func (r *Reader) At(ch uint, offset uint64) (float32, error) {
	base := r.data + offset*uint64(r.chans)*uint64(r.bits)/8
	base += uint64(ch) * uint64(r.bits) / 8
	switch r.bits {
	case 8:
		u8, err := readU8(r.r, int64(base))
		if err != nil {
			return 0.0, err
		}
		return float32(u8) / 256, nil

	case 16:
		s16, err := readS16(r.r, int64(base))
		if err != nil {
			return 0.0, err
		}
		return float32(s16) / (1 << 15), nil
	}
	return 0.0, fmt.Errorf("unsupported bit size")
}

// Slice ...
func (r *Reader) Slice(sl, sr uint64) (*Reader, error) {
	if sl >= r.samples {
		return nil, fmt.Errorf("sl >= samples")
	}
	if sl >= r.samples {
		return nil, fmt.Errorf("sl >= samples")
	}
	if sl >= sr {
		return nil, fmt.Errorf("sl >= sr")
	}
	base := r.data + sl*uint64(r.chans)*uint64(r.bits)/8
	return &Reader{
		r:       r.r,
		format:  r.format,
		chans:   r.chans,
		size:    r.size,
		bits:    r.bits,
		rate:    r.rate,
		data:    base,
		samples: sr - sl,
	}, nil
}
