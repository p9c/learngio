package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/kiktomo/goqr"
	"image"
)

func main() {
	//th := material.NewTheme()
	txt := "HELLO 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	version := 3

	// QRcode error correction level: "M"
	//eclevel := 0 (Auto)
	eclevel := goqr.ECLevelM
	// Encode() returns image.Image
	qr, err := goqr.Encode(txt, version, eclevel)
	if err != nil {
	}
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
						addrQR := paint.NewImageOp(qr)
						sz := gtx.Constraints.Width.Constrain(gtx.Px(unit.Dp(500)))
						addrQR.Add(gtx.Ops)
						paint.PaintOp{
							Rect: f32.Rectangle{
								Max: f32.Point{
									X: float32(sz), Y: float32(sz),
								},
							},
						}.Add(gtx.Ops)
						gtx.Dimensions.Size = image.Point{X: sz, Y: sz}
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
