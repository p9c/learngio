package main

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/helpers"
	"image"
)

type Button struct {
	pressed        bool
	Name           string
	ColorBg        string
	ColorBgHover   string
	ColorText      string
	ColorTextHover string
	BorderRadius   [4]float32
}

func button(name string) *Button {
	return &Button{
		Name:           name,
		ColorBg:        "ff303030",
		ColorBgHover:   "ffcf3030",
		ColorText:      "ffcfcfcf",
		ColorTextHover: "ffcfcfcf",
	}
}

func (b *Button) Layout(gtx *layout.Context, do func()) func() {
	return func() {
		for _, e := range gtx.Events(b) { // HLevent
			if e, ok := e.(pointer.Event); ok { // HLevent
				switch e.Type { // HLevent
				case pointer.Press: // HLevent
					b.pressed = true // HLevent
					do()
				case pointer.Release: // HLevent
					b.pressed = false // HLevent
				}
			}
		}

		cs := gtx.Constraints
		th := material.NewTheme()

		//colorBg := helpers.HexARGB("ff30cfcf")
		colorText := helpers.HexARGB(b.ColorText)
		colorBg := helpers.HexARGB(b.ColorBg)

		if b.pressed {
			colorText = helpers.HexARGB(b.ColorTextHover)
			colorBg = helpers.HexARGB(b.ColorBgHover)
		}
		pointer.Rect( // HLevent
			image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
		).Add(gtx.Ops) // HLevent
		pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))

		in := layout.UniformInset(unit.Dp(5))
		in.Layout(gtx, func() {
			helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))
			//cs := gtx.Constraints
			txt := th.Caption(b.Name)
			txt.Color = colorText
			txt.Layout(gtx)
		})
	}
}
