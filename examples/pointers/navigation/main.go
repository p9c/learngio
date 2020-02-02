package main

import (
	"fmt"
	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/font/gofont"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/unit"
	"github.com/gop9/olt/gio/widget/material"
	"github.com/p9c/learngio/helpers"
)

var (
	content = "content1"
	mainNav = &layout.List{
		Axis: layout.Vertical,
	}
)

func main() {
	go func() {
		w := app.NewWindow()
		gofont.Register()
		th := material.NewTheme()

		menuItem1 := &Button{
			Name: "Content 1",
			Do: func(n interface{}) {
				content = "content1"
			},
			OperateValue: 1,
		}
		menuItem2 := &Button{
			Name: "Content 2",
			Do: func(n interface{}) {
				content = "content2"
			},
			OperateValue: 1,
		}
		menuItem3 := &Button{
			Name: "Content 3",
			Do: func(interface{}) {
				content = "content3"
			},
			OperateValue: 0,
		}
		menuItem4 := &Button{
			Name: "Content 4",
			Do: func(interface{}) {
				content = "content4"
			},
			OperateValue: 0,
		}

		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func() {

						//navButtons := []func(){
						//	func() {
						//		menuItem1.Layout(gtx)
						//	},
						//	func() {
						//		menuItem2.Layout(gtx)
						//	},
						//	func() {
						//		menuItem3.Layout(gtx)
						//	},
						//}
						//mainNav.Layout(gtx, len(navButtons), func(i int) {
						//	layout.UniformInset(unit.Dp(0)).Layout(gtx, navButtons[i])
						//})

						layout.Flex{
							Axis: layout.Vertical,
						}.Layout(gtx,
							layout.Flexed(0.25, func() {
								menuItem1.Layout(gtx)
							}),
							layout.Flexed(0.25, func() {
								menuItem2.Layout(gtx)
							}),
							layout.Flexed(0.25, func() {
								menuItem3.Layout(gtx)
							}),
							layout.Flexed(0.25, func() {
								menuItem4.Layout(gtx)
							}),
						)

					}),

					layout.Flexed(1, func() {

						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

						in := layout.UniformInset(unit.Dp(0))
						in.Layout(gtx, func() {
							layout.Align(layout.Center).Layout(gtx, func() {
								th.H3(fmt.Sprint(content)).Layout(gtx)
							})
						})
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
