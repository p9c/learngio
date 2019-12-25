package simple_pointers

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/components/button"
	"github.com/p9c/learngio/helpers"
	"image/color"
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

		increase := &button.Button{
			Name: "increase",
			Do: func(interface{}) {
				itemValue.doIncrease(1)
			},
		}
		decrease := &button.Button{
			Name: "decrease",
			Do: func(interface{}) {
				itemValue.doDecrease(1)
			},
		}
		reset := &button.Button{
			Name: "reset",
			Do: func(interface{}) {
				itemValue.doReset()
			},
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
										helpers.DrawRectangle(gtx, cs.Width.Max, 120, color.RGBA{A: 0xff, R: 0x30, G: 0x30, B: 0xcf}, 0, 0, 0, 0, unit.Dp(0))

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
