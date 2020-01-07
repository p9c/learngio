package helpers

import (
	"image"
	"image/color"

	"github.com/p9c/learngio/pkg/gio/f32"
	"github.com/p9c/learngio/pkg/gio/layout"
	"github.com/p9c/learngio/pkg/gio/op/clip"
	"github.com/p9c/learngio/pkg/gio/op/paint"
	"github.com/p9c/learngio/pkg/gio/unit"
)

func DrawRectangle(gtx *layout.Context, w, h int, color color.RGBA, borderRadius [4]float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		//cs := gtx.Constraints
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(w),
				Y: float32(h),
			},
		}
		paint.ColorOp{Color: color}.Add(gtx.Ops)

		clip.Rect{Rect: square,
			NE: borderRadius[0], NW: borderRadius[1], SE: borderRadius[2], SW: borderRadius[3]}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: w, Y: h}}
	})
}
