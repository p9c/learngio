package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

var (
	list = &layout.List{
		Axis:        layout.Horizontal,
		ScrollToEnd: true,
	}
)

func main() {

	gofont.Register()
	th := material.NewTheme()

	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(5), Right: unit.Dp(4)}.Layout(gtx, func() {

					drawList(gtx, list, th)

				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func drawList(gtx *layout.Context, list *layout.List, th *material.Theme) {
	n := []int{25, 15, 26, 21, 24, 26, 33, 25, 15, 25, 22, 24, 33, 25, 15, 25, 22, 24, 44, 45, 47, 49, 55, 66, 77, 88, 79, 75, 54, 3, 3, 3, 4, 5, 6, 7, 8, 9, 11, 22, 33, 44, 55, 66, 77, 88, 99, 25, 15, 26, 21, 24, 26, 33, 25, 15, 25, 22, 24, 33, 25, 15, 25, 22, 24, 44, 45, 47, 49, 55, 66, 77, 88, 79, 75, 54, 3, 3, 3, 4, 5, 6, 7, 8, 9, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	list.Layout(gtx, len(n), func(i int) {
		drawRectangle(gtx, color.RGBA{0x00, 0x00, 0xff, 0xff}, gtx.Constraints.Height.Max)
		drawRectangle(gtx, color.RGBA{0x00, 0x00, 0xff, 0xff}, n[i]+1)
		drawRectangle(gtx, color.RGBA{0xff, 0x00, 0x00, 0xff}, n[i])
	})
}

func drawRectangle(gtx *layout.Context, color color.RGBA, value int) {
	in := layout.UniformInset(unit.Dp(0))
	in.Layout(gtx, func() {
		//cs := gtx.Constraints
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(9),
				Y: float32(value),
			},
		}
		paint.ColorOp{Color: color}.Add(gtx.Ops)

		clip.Rect{Rect: square}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: 9, Y: value}}
	})
}
