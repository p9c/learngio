package main

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"

	"github.com/p9c/learngio/helpers"
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{
					Spacing: layout.SpaceBetween,
				}.Layout(gtx,
					layout.Flexed(0.25, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0x30}, 0, 0, 0, 0)
					}),
					layout.Flexed(0.25, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, color.RGBA{A: 0xff, R: 0x30, G: 0x30, B: 0xcf}, 0, 0, 0, 0)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
