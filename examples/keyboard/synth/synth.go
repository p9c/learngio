package main

import (
	"math"
	"sync"
)

const SampleRate = 44100

type Synth struct {
	Channels    []Channel
	NextChannel int

	Volume     float32
	Attack     float32
	Release    float32
	Modulation float32
	sync.Mutex
}

func NewSynth(channels int) *Synth {
	s := &Synth{Channels: make([]Channel, channels)}
	for i := range s.Channels {
		s.Channels[i].s = s
	}
	return s
}

func (s *Synth) NoteOn(p int) {
	s.Lock()
	c := &s.Channels[s.NextChannel]
	c.Pitch = p
	c.Frequency = 440 * float32(math.Pow(2, float64(p-69)/12)) / SampleRate
	c.On = true

	s.NextChannel++
	if s.NextChannel >= len(s.Channels) {
		s.NextChannel = 0
	}
	s.Unlock()
}

func (s *Synth) NoteOff(p int) {
	s.Lock()
	for i := range s.Channels {
		c := &s.Channels[i]
		if c.Pitch == p {
			c.On = false
		}
	}
	s.Unlock()
}

func (s *Synth) SetVolume(volume float32) {
	s.Lock()
	s.Volume = volume * .1
	s.Unlock()
}

func (s *Synth) SetAttack(attack float32) {
	s.Lock()
	s.Attack = 1 - float32(math.Pow(0.1, 1.0/(float64(attack)*SampleRate)))
	s.Unlock()
}

func (s *Synth) SetRelease(release float32) {
	s.Lock()
	s.Release = 1 - float32(math.Pow(0.1, 1.0/(float64(release)*SampleRate)))
	s.Unlock()
}

func (s *Synth) SetModulation(mod float32) {
	s.Lock()
	s.Modulation = mod
	s.Unlock()
}

func (s *Synth) Process(out []float32) {
	s.Lock()
	for i := range s.Channels {
		s.Channels[i].Process(out)
	}
	for i := range out {
		out[i] *= s.Volume
	}
	s.Unlock()
}

type Channel struct {
	Pitch     int
	Frequency float32
	On        bool

	s *Synth

	phase  float32
	env    float32
	mod    float32
	x1, x2 float32
	y1, y2 float32
}

func (c *Channel) Process(out []float32) {
	if c.Frequency == 0 {
		return
	}
	for i := range out {
		// calculate the phase
		c.phase += c.Frequency
		if c.phase >= 1 {
			c.phase--
		}

		// calculate the envelope
		if c.On {
			c.env += (1 - c.env) * c.s.Attack
		} else {
			c.env -= c.env * c.s.Release
		}

		// modulate the pulse width
		c.mod += 4.0 / SampleRate
		if c.mod >= 3 {
			c.mod -= 4
		}
		w := c.mod
		if w > 1 {
			w = 2 - w
		}
		w *= c.s.Modulation / 2

		// calculate the waveform
		x := sqr(c.phase, 0.5+w, c.Frequency) * c.env

		// apply a low pass filter
		y := 0.001995109*x + 0.003990218*c.x1 + 0.001995109*c.x2 + 1.9334036*c.y1 - 0.9413841*c.y2
		c.x1, c.x2, c.y1, c.y2 = x, c.x1, y, c.y1

		out[i] += y
	}
}

func saw(p, f float32) float32 {
	q := p + f
	r := q + f
	if r > 1 {
		r--
		if q > 1 {
			q--
		}
	}
	return (((2*p-3)*p+1)*p - ((4*q-6)*q+2)*q + ((2*r-3)*r+1)*r) / (6 * f * f)
}

func sqr(p, w, f float32) float32 {
	q := p + w
	if q > 1 {
		q--
	}
	return saw(p, f) - saw(q, f)
}
