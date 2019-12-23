package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/p9c/learngio/helpers"
	"image/color"
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				drawRects(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(gtx *layout.Context) {
	cs := gtx.Constraints
	stack := layout.Stack{Alignment: layout.Center}

	stack.Layout(gtx,
		layout.Stacked(func() {
			helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0x30}, 0, 0, 0, 0, unit.Dp(64))

		}),
	)

	stack.Layout(gtx,
		layout.Stacked(func() {
			helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0xcf}, 0, 0, 0, 0, unit.Dp(128))
		}),
	)

	stack.Layout(gtx,
		layout.Stacked(func() {
			helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0xcf}, 0, 0, 0, 0, unit.Dp(256))

		}),
	)

}

// END OMIT
