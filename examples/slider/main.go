package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func main() {
	//itemValue := item{
	//	i: 0,
	//}
	panel := onePanel()
	go func() {
		w := app.NewWindow()
		gofont.Register()
		th := material.NewTheme()
		//sl := slider(th, 17)
		gtx := layout.NewContext(w.Queue())
		panel.listObjects = listObjects(gtx, th)
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				panel.Layout(gtx, th)
				//fmt.Println(panel.listHeight)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
