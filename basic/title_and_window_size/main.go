package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/p9c/learngio/helpers"
)

func main() {

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
					layout.Flexed(0.5, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ffcf30cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					}),
					layout.Flexed(0.25, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					}),
					layout.Flexed(0.25, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
