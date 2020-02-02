package main

import (
	"fmt"

	"github.com/jfreymuth/ui"
	"github.com/jfreymuth/ui/draw"
	"github.com/jfreymuth/ui/text"
)

type Slider struct {
	Label string
	// two Text structs, one for the label and one for the value
	label text.Text
	text  text.Text
	value float32

	// Usually a callback's signature should include *ui.State,
	// we don't do that here because we only want to call Synth.Set...()
	Changed func(float32)
}

func NewSlider(label string, init float32, changed func(float32)) *Slider {
	if changed != nil {
		changed(init)
	}
	return &Slider{Label: label, value: init, Changed: changed}
}

func (s *Slider) PreferredSize(fonts draw.FontLookup) (int, int) {
	w, _ := s.label.Size(s.Label, draw.Font{Name: "default", Size: 12}, fonts)
	return w + 10, 200
}

func (s *Slider) Update(g *draw.Buffer, state *ui.State) {
	w, h := g.Size()
	// determine the height of the text
	_, lh := s.label.Size(s.Label, draw.Font{Name: "default", Size: 12}, g.FontLookup)
	// calculate the lengh of the actual slider
	l := h - 20 - 2*lh

	// handle input
	if state.MouseButtonDown(ui.MouseLeft) {
		my := state.MousePos().Y
		v := float32(my-10-lh) / float32(l)
		if v < 0 {
			v = 0
		} else if v > 1 {
			v = 1
		}
		if s.value != v {
			s.value = v
			if s.Changed != nil {
				s.Changed(v)
			}
		}
	}

	// draw text
	s.label.DrawCentered(g, draw.XYXY(0, 0, w, lh), s.Label, draw.Font{Name: "default", Size: 12}, draw.Black)
	s.text.DrawCentered(g, draw.XYXY(0, h-lh, w, h), fmt.Sprintf("%.2f", s.value), draw.Font{Name: "default", Size: 12}, draw.Black)

	// draw the track
	x := w / 2
	g.Fill(draw.XYWH(x-1, 10+lh, 2, l), draw.RGBA(0, 0, 0, .4))
	// draw the handle
	y := 10 + lh + int(s.value*float32(l))
	g.Shadow(draw.XYWH(x-5, y-5, 10, 10), draw.Black, 5)
	g.Fill(draw.XYWH(x-5, y-5, 10, 10), draw.White)
}
