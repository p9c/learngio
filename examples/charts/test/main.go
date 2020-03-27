package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/fogleman/gg"
	"image"
	"math"
)

type Point struct {
	X, Y float64
}

func Polygon(n int, x, y, r float64) []Point {
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
	}
	return result
}

func main() {
	//th := material.NewTheme()
	// Initialize the graphic context on an RGBA image

	n := 56
	points := Polygon(n, 1512, 1512, 400)
	dc := gg.NewContext(2024, 2024)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()

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
						addrQR := paint.NewImageOp(dc.Image())
						sz := gtx.Constraints.Width.Constrain(gtx.Px(unit.Dp(100)))
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
