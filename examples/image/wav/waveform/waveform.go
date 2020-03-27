package waveform

import (
	"image"
	"image/color"
	"log"
	"math"
)

// WaveReader interface ...
type WaveReader interface {
	Len() uint64
	Rate() uint32
	Chans() uint16
	At(ch uint, offset uint64) (float32, error)
}

// Options represents ...
type Options struct {
	Width   int
	Height  int
	Half    bool
	Zoom    float32
	MarginL int
	MarginR int
	MarginT int
	MarginB int
	Back    *color.NRGBA
	Front   *color.NRGBA
}

func initOptions(o *Options) *Options {
	no := &Options{
		Width:  800,
		Height: 250,
		Zoom:   0.8,
		Front: &color.NRGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Back: &color.NRGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		},
	}
	if o != nil {
		if o.Half {
			no.Half = true
		}
		if o.Width > 0 {
			no.Width = o.Width
		}
		if o.Height > 0 {
			no.Height = o.Height
		}
		if o.Back != nil {
			no.Back = o.Back
		}
		if o.Front != nil {
			no.Front = o.Front
		}
		if o.Zoom != 0 {
			no.Zoom = o.Zoom
		}
		if o.MarginL >= 0 {
			no.MarginL = o.MarginL
		}
		if o.MarginR >= 0 {
			no.MarginR = o.MarginR
		}
		if o.MarginT >= 0 {
			no.MarginT = o.MarginT
		}
		if o.MarginB >= 0 {
			no.MarginB = o.MarginB
		}
	}
	return no
}

func newNRGBA(o *Options) *image.NRGBA {
	if o == nil {
		panic("options 'o' is nil")
	}
	dx := o.Width + o.MarginL + o.MarginR
	dy := o.Height + o.MarginT + o.MarginB
	rc := image.Rect(0, 0, dx, dy)
	im := image.NewNRGBA(rc)
	for y := 0; y < dx; y++ {
		for x := 0; x < dy; x++ {
			im.SetNRGBA(x, y, *o.Back)
		}
	}
	return im
}

func getMinMax(w WaveReader, pos, size int) (float32, float32) {
	s := float64(w.Len())
	f := float64(pos) / float64(size)
	t := float64(pos+1) / float64(size)

	fo := uint64(f * s)
	to := uint64(t * s)
	min := float32(+1.0)
	max := float32(-1.0)
	for o := fo; o < to; o++ {
		h, err := w.At(0, o)
		if err != nil {
			log.Fatalf("w.At %d failed: %v", o, err)
		}
		if h < min {
			min = h
		}
		if h > max {
			max = h
		}
	}
	return min, max
}

// MinMax function ...
func MinMax(w WaveReader, o *Options) *image.NRGBA {
	o = initOptions(o)
	wf := newNRGBA(o)
	if uint64(o.Width) < w.Len() {
		for x := 0; x < o.Width; x++ {
			h0, h1 := getMinMax(w, x, o.Width)
			m := float32(o.Height) / 2
			H0 := h0 * m // must be negative
			H1 := h1 * m
			t := m - H1*o.Zoom
			b := m - H0*o.Zoom
			for y := int(t); y < int(b); y++ {
				wf.SetNRGBA(x+o.MarginL, y+o.MarginT, *o.Front)
			}
		}
		return wf
	}
	return wf
}

func getAbsMax(w WaveReader, pos, size int) float32 {
	s := float64(w.Len())
	f := float64(pos) / float64(size)
	t := float64(pos+1) / float64(size)

	fo := uint64(f * s)
	to := uint64(t * s)
	m := float32(0.0)
	for o := fo; o < to; o++ {
		h, err := w.At(0, o)
		if err != nil {
			log.Fatalf("w.At %d failed: %v", o, err)
		}
		if h < 0 {
			h = -h
		}
		if h > m {
			m = h
		}
	}
	return m
}

// AbsMax function ...
func AbsMax(w WaveReader, o *Options) *image.NRGBA {
	o = initOptions(o)
	wf := newNRGBA(o)
	if uint64(o.Width) < w.Len() {
		for x := 0; x < o.Width; x++ {
			h := getAbsMax(w, x, o.Width)
			m := float32(o.Height) / 2
			H := h * m
			t := m - H*o.Zoom
			b := m + H*o.Zoom
			if o.Half {
				b = float32(o.Height)
				t = b - h*b*o.Zoom
			}
			for y := int(t); y < int(b); y++ {
				wf.SetNRGBA(x+o.MarginL, y+o.MarginT, *o.Front)
			}
		}
		return wf
	}
	return wf
}

func getRms(w WaveReader, pos, size int) float32 {
	s := float64(w.Len())
	f := float64(pos) / float64(size)
	t := float64(pos+1) / float64(size)

	fo := uint64(f * s)
	to := uint64(t * s)
	ss := float32(0.0)
	for o := fo; o < to; o++ {
		h, err := w.At(0, o)
		if err != nil {
			log.Fatalf("w.At %d failed: %v", o, err)
		}
		ss += h * h
	}
	ss /= float32(to - fo)
	sq := math.Sqrt(float64(ss))
	return float32(sq)
}

// Rms function ...
func Rms(w WaveReader, o *Options) *image.NRGBA {
	o = initOptions(o)
	wf := newNRGBA(o)
	if uint64(o.Width) < w.Len() {
		for x := 0; x < o.Width; x++ {
			h := getRms(w, x, o.Width)
			m := float32(o.Height) / 2
			H := h * m
			t := m - H*o.Zoom
			b := m + H*o.Zoom
			if o.Half {
				b = float32(o.Height)
				t = b - h*b*o.Zoom
			}
			for y := int(t); y < int(b); y++ {
				wf.SetNRGBA(x+o.MarginL, y+o.MarginT, *o.Front)
			}
		}
		return wf
	}
	return wf
}
