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

func (it *item) doSlide(n int) {
	it.i = it.i + int(n)
}

func main() {
	itemValue := item{
		i: 0,
	}
	go func() {
		w := app.NewWindow()
		gofont.Register()
		th := material.NewTheme()

		slider := &Slider{
			Name: "increase",
			Do: func(n interface{}) {
				itemValue.doSlide(n.(int))
			},
			OperateValue: 1,
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
									layout.Rigid(func() {
										cs := gtx.Constraints
										helpers.DrawRectangle(gtx, cs.Width.Max, 120, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

										in := layout.UniformInset(unit.Dp(0))
										in.Layout(gtx, func() {
											layout.Align(layout.Center).Layout(gtx, func() {
												th.H3(fmt.Sprint(slider.Position)).Layout(gtx)
											})
										})
									}),
									layout.Flexed(1, func() {
										layout.Flex{}.Layout(gtx,
											layout.Flexed(0.4, func() {
												slider.Layout(gtx)
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

type Slider struct {
	pressed      bool
	Name         string
	Do           func(interface{})
	ColorBg      string
	Position     float32
	BorderRadius [4]float32
	OperateValue interface{}
}

func (s *Slider) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(s) { // HLevent

		if e, ok := e.(pointer.Event); ok { // HLevent
			fmt.Println("Event: ", e.Position.Y)
			s.Position = e.Position.Y
			fmt.Println("**********************************************************")
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				s.pressed = true // HLevent
				s.Do(s.OperateValue)
			case pointer.Release: // HLevent
				s.pressed = false // HLevent
			}
		}
	}

	cs := gtx.Constraints
	//th := material.NewTheme()

	sliderBg := helpers.HexARGB("ff558899")
	colorBg := helpers.HexARGB("ff30cfcf")
	colorBorder := helpers.HexARGB("ffcf3030")
	border := unit.Dp(0)
	var position float32 = 0
	if s.pressed {
		position = s.Position
		colorBg = helpers.HexARGB("ffcf30cf")
		colorBorder = helpers.HexARGB("ff303030")
		border = unit.Dp(3)
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: s}.Add(gtx.Ops) // HLevent
	helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBorder, s.BorderRadius, unit.Dp(0))

	in := layout.UniformInset(border)
	in.Layout(gtx, func() {
		cs := gtx.Constraints
		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, s.BorderRadius, unit.Dp(0))
		//cs := gtx.Constraints
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func() {
				layout.Align(layout.Center).Layout(gtx, func() {

					in := layout.Inset{
						Top: unit.Dp(position),
					}
					in.Layout(gtx, func() {
						cs := gtx.Constraints

						helpers.DrawRectangle(gtx, cs.Width.Max, 50, sliderBg, [4]float32{15, 15, 15, 15}, unit.Dp(0))
					})

				})
			}),
		)
	})
}
