package framework

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/helpers"
	"image"
	"strconv"
)

type Input struct {
	pressed bool
	Name    string
	Do      func(interface{})

	ColorBg      string
	OperateValue interface{}
}

var (
	lineEditor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func (b *Input) Layout(gtx *layout.Context) {
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

	col := helpers.HexARGB("ff30cfcf")
	if b.pressed {
		col = helpers.HexARGB("ffcf30cf")
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, col, [4]float32{0, 0, 0, 0}, unit.Dp(0))

	layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(0.4, func() {
			layout.Align(layout.Center).Layout(gtx, func() {
				th.H5(b.Name).Layout(gtx)
			})

		}),
		layout.Flexed(0.2, func() {
			//reset.Layout(gtx)

			e := th.Editor("Hint")
			e.Font.Style = text.Italic
			e.Layout(gtx, lineEditor)
			for _, e := range lineEditor.Events(gtx) {
				if e, ok := e.(widget.SubmitEvent); ok {
					b.OperateValue, _ = strconv.Atoi(e.Text)
					lineEditor.SetText("")
				}
			}

		}),
		layout.Flexed(0.4, func() {

			layout.Align(layout.Center).Layout(gtx, func() {
				th.H5(b.Name).Layout(gtx)
			})

		}),
	)

}
