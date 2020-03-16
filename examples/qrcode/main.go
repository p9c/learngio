package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/nfnt/resize"
	"image"
)

func main() {
	qr, err := Encode("HELLO 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 0, ECLevelM)
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
						qrResize := resize.Resize(555, 0, qr, resize.NearestNeighbor)
						addrQR := paint.NewImageOp(qrResize)

						scale := float32(160 / 72)

						size := addrQR.Rect.Size()
						wf, hf := float32(size.X), float32(size.Y)
						w, h := gtx.Px(unit.Dp(wf*scale)), gtx.Px(unit.Dp(hf*scale))
						cs := gtx.Constraints
						d := image.Point{X: cs.Width.Constrain(w), Y: cs.Height.Constrain(h)}
						var s op.StackOp
						s.Push(gtx.Ops)
						clip.Rect{Rect: f32.Rectangle{Max: toPointF(d)}}.Op(gtx.Ops).Add(gtx.Ops)
						addrQR.Add(gtx.Ops)
						paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(w), Y: float32(h)}}}.Add(gtx.Ops)
						s.Pop()
						gtx.Dimensions = layout.Dimensions{Size: d}
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
