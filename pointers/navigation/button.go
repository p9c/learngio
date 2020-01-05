package main

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
	pressed           bool
	Name              string
	Do                func(interface{})
	ColorBg           string
	BorderRadius      [4]float32
	OperateValue      interface{}
	PaddingVertical   unit.Value
	PaddingHorizontal unit.Value
	Icon              *Icon
	IconSize          int
	IconColor         color.RGBA
}

func (b *Button) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
				b.Do(b.OperateValue)
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}

	cs := gtx.Constraints
	th := material.NewTheme()

	colorBg := helpers.HexARGB("ff30cfcf")
	colorBorder := helpers.HexARGB("ffcf3030")
	border := unit.Dp(1)

	if b.pressed {
		colorBg = helpers.HexARGB("ffcf30cf")
		colorBorder = helpers.HexARGB("ff303030")
		border = unit.Dp(3)
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBorder, b.BorderRadius, unit.Dp(0))

	in := layout.UniformInset(border)
	in.Layout(gtx, func() {
		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))
		//cs := gtx.Constraints
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.3, func() {
				layout.Align(layout.Center).Layout(gtx, func() {
					th.H6(b.Name).Layout(gtx)
				})
			}),
			layout.Flexed(0.5, func() {
				//reset.Layout(gtx)
				layout.Align(layout.Center).Layout(gtx, func() {
					th.H5(b.Name).Layout(gtx)
				})
			}),
			layout.Flexed(0.2, func() {
				layout.Align(layout.Center).Layout(gtx, func() {
					th.Caption(b.Name).Layout(gtx)
				})
			}),
		)
	})
}
