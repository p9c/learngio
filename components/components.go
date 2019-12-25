package components

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/helpers"
	"image"
	"image/color"
)

type Button struct {
	pressed bool
	Name    string
	Do      func()
}

func (b *Button) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				//b.pressed = true // HLevent
				b.Do()
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}

	cs := gtx.Constraints
	th := material.NewTheme()

	col := color.RGBA{A: 0xff, R: 0xff}
	if b.pressed {
		col = color.RGBA{A: 0xff, G: 0xff}
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, col, 0, 0, 0, 0, unit.Dp(0))
	layout.Align(layout.Center).Layout(gtx, func() {
		th.H3(b.Name).Layout(gtx)
	})
}
