package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
)

var (
	//switchItem = new(cButton)
	switchItem = new(CheckBox)
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

				//sw := NewButton()
				//sw.Layout(gtx, switchItem)
				sw := NewDuoUIcheckBox()

				sw.Layout(gtx, switchItem)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
