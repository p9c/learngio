package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func main() {
	gofont.Register()
	th := material.NewTheme()
	panel := onePanel(th)
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				panel.Layout(gtx, testContent(gtx, th))
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
