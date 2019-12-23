package helpers

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func DrawRectangle(gtx *layout.Context, w, h int, color color.RGBA, ne, nw, se, sw float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		cs := gtx.Constraints
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(cs.Width.Max),
				Y: float32(cs.Height.Max),
			},
		}
		paint.ColorOp{Color: color}.Add(gtx.Ops)

		clip.Rect{Rect: square,
			NE: ne, NW: nw, SE: se, SW: sw}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: w, Y: h}}
	})
}
