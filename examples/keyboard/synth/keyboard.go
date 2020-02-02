package main

import (
	"github.com/jfreymuth/ui"
	"github.com/jfreymuth/ui/draw"
)

type Keyboard struct {
	NoteOn, NoteOff func(int)
	KeyWidth        int
	last            int
	sustain         bool
	notesOn         [2]uint64
}

func NewKeyboard() *Keyboard {
	return &Keyboard{KeyWidth: 36}
}

func (k *Keyboard) PreferredSize(fonts draw.FontLookup) (int, int) {
	return (6*7 + 3) * k.KeyWidth, 160
}

func (k *Keyboard) Update(g *draw.Buffer, state *ui.State) {
	pitch := 0
	hovered := 0
	if state.HasMouseFocus() {
		mouse := state.MousePos()
		hovered = k.getPitch(mouse.X+12*k.KeyWidth, mouse.Y)
		if hovered < 21 || hovered > 96 {
			hovered = 0
		}
		if state.MouseButtonDown(ui.MouseLeft) {
			pitch = hovered
		}
	}
	if k.last != pitch {
		if k.last != 0 && !state.HasModifiers(ui.Shift) {
			k.noteOff(k.last)
		}
		if pitch != 0 {
			k.noteOn(pitch)
			k.last = pitch
		}
	}
	if k.sustain && !state.HasModifiers(ui.Shift) {
		k.releaseSustain(pitch)
	}
	k.last = pitch
	k.sustain = state.HasModifiers(ui.Shift)
	k.draw(g, hovered)
}

func (k *Keyboard) draw(g *draw.Buffer, hovered int) {
	kw := k.KeyWidth
	k.drawKeyW(g, 0, hovered == 21)
	k.drawKeyW(g, kw, hovered == 23)
	k.drawKeyB(g, 6*kw/7, hovered == 22)
	for o := 0; o < 6; o++ {
		x := (o*7 + 2) * kw
		for i := 0; i < 3; i++ {
			k.drawKeyW(g, x+i*kw, hovered == (o+2)*12+i*2)
		}
		for i := 3; i < 7; i++ {
			k.drawKeyW(g, x+i*kw, hovered == (o+2)*12+i*2-1)
		}

		k.drawKeyB(g, x+3*kw/5, hovered == (o+2)*12+1)
		k.drawKeyB(g, x+9*kw/5, hovered == (o+2)*12+3)
		k.drawKeyB(g, x+25*kw/7, hovered == (o+2)*12+6)
		k.drawKeyB(g, x+33*kw/7, hovered == (o+2)*12+8)
		k.drawKeyB(g, x+41*kw/7, hovered == (o+2)*12+10)
	}
	k.drawKeyW(g, 44*kw, hovered == 96)
}

func (k *Keyboard) drawKeyW(g *draw.Buffer, x int, h bool) {
	g.Outline(draw.XYWH(x, 0, k.KeyWidth, 160), draw.Black)
	if h {
		g.Fill(draw.XYWH(x+1, 1, k.KeyWidth-2, 158), draw.Gray(.9))
	}
}

func (k *Keyboard) drawKeyB(g *draw.Buffer, x int, h bool) {
	if h {
		g.Fill(draw.XYWH(x, 0, k.KeyWidth*3/5, 100), draw.Gray(.3))
	} else {
		g.Fill(draw.XYWH(x, 0, k.KeyWidth*3/5, 100), draw.Black)
	}
}

func (k *Keyboard) getPitch(x, y int) int {
	kw := k.KeyWidth
	octave := x / (7 * kw)
	x -= octave * 7 * kw
	if y < 100 {
		if x < 3*kw {
			return octave*12 + x*5/(3*kw)
		} else {
			return octave*12 + 5 + (x-3*kw)*7/(4*kw)
		}
	}
	w := x / kw
	if w < 3 {
		return octave*12 + w*2
	} else {
		return octave*12 + w*2 - 1
	}
}

func (k *Keyboard) noteOn(p int) {
	k.notesOn[p>>6] |= 1 << uint(p&63)
	if k.NoteOn != nil {
		k.NoteOn(p)
	}
}

func (k *Keyboard) noteOff(p int) {
	k.notesOn[p>>6] &^= 1 << uint(p&63)
	if k.NoteOff != nil {
		k.NoteOff(p)
	}
}

func (k *Keyboard) releaseSustain(p int) {
	for i := 0; i < 128; i++ {
		if i != p && k.notesOn[i>>6]&(1<<uint(i&63)) != 0 {
			k.noteOff(i)
		}
	}
}
