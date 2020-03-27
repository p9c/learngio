package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
	"image/color"
)

func main() {
	//th := material.NewTheme()
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {

			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Flexed(1, func() {
						drawGradient(gtx)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func drawGradient(gtx *layout.Context) {
	width := gtx.Constraints.Width.Max
	height := gtx.Constraints.Height.Max
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	size := m.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
				uint8(48 * x / size.X),
				uint8(255 * y / size.Y),
				255,
				255}
			m.Set(x, y, color)
		}
	}
	addrQR := paint.NewImageOp(m)
	//sz := gtx.Constraints.Width.Constrain(gtx.Px(unit.Dp(100)))
	addrQR.Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: f32.Point{
				X: float32(width), Y: float32(height),
			},
		},
	}.Add(gtx.Ops)
	gtx.Dimensions.Size = image.Point{X: width, Y: height}
}
