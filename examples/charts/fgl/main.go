package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/wrnrlr/shape"
	"image/color"
)

var (
	list = &layout.List{
		Axis:        layout.Horizontal,
		ScrollToEnd: true,
	}
)

const numPoints = 23

func main() {

	gofont.Register()

	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Center.Layout(gtx, func() {

					drawLine(gtx)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func drawLine(gtx *layout.Context) {
	col := color.RGBA{255, 0, 0, 255}
	colo := color.RGBA{0, 255, 0, 255}

	fff := f32.Point{
		X: 0,
		Y: 0,
	}
	shape.Circle{fff, 100}.Fill(colo, gtx)
	shape.Circle{fff, 100}.Stroke(col, 0, gtx)

}
