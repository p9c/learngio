package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/helpers"
	"image"
)

type item struct {
	i int
}

func (it *item) doIncrease(n int) {
	it.i = it.i + int(n)
}

func (it *item) doDecrease(n int) {
	it.i = it.i - int(n)
}
func (it *item) doReset() {
	it.i = 0
}

func main() {
	itemValue := item{
		i: 0,
	}
	go func() {
		w := app.NewWindow()
		gofont.Register()
		th := material.NewTheme()

		increase := &Button{
			Name: "increase",
			Do: func(n interface{}) {
				itemValue.doIncrease(n.(int))
			},
			OperateValue: 1,
		}
		decrease := &Button{
			Name: "decrease",
			Do: func(n interface{}) {
				itemValue.doDecrease(n.(int))
			},
			OperateValue: 1,
		}
		reset := &Button{
			Name: "reset",
			Do: func(interface{}) {
				itemValue.doReset()
			},
			OperateValue: 0,
		}

		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{
					Axis:    layout.Horizontal,
					Spacing: layout.SpaceSides,
				}.Layout(gtx,
					layout.Flexed(0.5, func() {
						layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceSides,
						}.Layout(gtx,
							layout.Flexed(0.5, func() {

								layout.Flex{Axis: layout.Vertical}.Layout(gtx,
									layout.Rigid(func() {
										cs := gtx.Constraints
										helpers.DrawRectangle(gtx, cs.Width.Max, 120, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

										in := layout.UniformInset(unit.Dp(0))
										in.Layout(gtx, func() {
											layout.Align(layout.Center).Layout(gtx, func() {
												th.H3(fmt.Sprint(itemValue.i)).Layout(gtx)
											})
										})
									}),
									layout.Flexed(1, func() {
										layout.Flex{}.Layout(gtx,
											layout.Flexed(0.4, func() {
												increase.Layout(gtx)
											}),
											layout.Flexed(0.2, func() {
												reset.Layout(gtx)
											}),
											layout.Flexed(0.4, func() {
												decrease.Layout(gtx)
											}),
										)
									}),
								)

							}),
						)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

type Button struct {
	pressed      bool
	Name         string
	Do           func(interface{})
	ColorBg      string
	BorderRadius [4]float32
	OperateValue interface{}
}

func (b *Button) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
				b.Do(b.OperateValue)
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}

	cs := gtx.Constraints
	th := material.NewTheme()

	colorBg := helpers.HexARGB("ff30cfcf")
	colorBorder := helpers.HexARGB("ffcf3030")
	border := unit.Dp(1)

	if b.pressed {
		colorBg = helpers.HexARGB("ffcf30cf")
		colorBorder = helpers.HexARGB("ff303030")
		border = unit.Dp(3)
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBorder, b.BorderRadius, unit.Dp(0))

	in := layout.UniformInset(border)
	in.Layout(gtx, func() {
		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))
		//cs := gtx.Constraints
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.3, func() {
				layout.Align(layout.Center).Layout(gtx, func() {
					th.H6(b.Name).Layout(gtx)
				})
			}),
			layout.Flexed(0.5, func() {
				//reset.Layout(gtx)
				layout.Align(layout.Center).Layout(gtx, func() {
					th.H5(b.Name).Layout(gtx)
				})
			}),
			layout.Flexed(0.2, func() {
				layout.Align(layout.Center).Layout(gtx, func() {
					th.Caption(b.Name).Layout(gtx)
				})
			}),
		)
	})
}
