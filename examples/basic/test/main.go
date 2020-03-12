package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	gofont.Register()
	th := material.NewTheme()
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func() {
						th.H6("header").Layout(gtx)
					}),

					layout.Flexed(1, func() {

						layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Rigid(func() {
								th.H6("sidebar left").Layout(gtx)
							}),

							layout.Flexed(0.5, func() {

								th.H6("content left").Layout(gtx)
							}),
							layout.Flexed(0.5, func() {

								th.H6("content right").Layout(gtx)
							}),

							layout.Rigid(func() {
								th.H6("sidebar right").Layout(gtx)
							}),
						)

					}),

					layout.Rigid(func() {
						th.H6("footer").Layout(gtx)
					}),
				)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
