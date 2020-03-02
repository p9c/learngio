package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

type button struct {
	pressed      bool
	Name         string
	Do           func(interface{})
	ColorBg      string
	BorderRadius [4]float32
	OperateValue interface{}
}

func (b *button) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(b) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				b.pressed = true
				b.Do(b.OperateValue)
			case pointer.Release:
				b.pressed = false
			}
		}
	}

	cs := gtx.Constraints
	th := material.NewTheme()
	if b.pressed {
	}
	pointer.Rect(
		image.Rectangle{Max: image.Point{X: cs.Width.Min, Y: cs.Height.Min}},
	).Add(gtx.Ops)
	pointer.InputOp{Key: b}.Add(gtx.Ops)
	th.H6(b.Name).Layout(gtx)
}

type DuoUIcounter struct {
	increase *button
	decrease *button
	reset    *button
	item     *item
	Font     text.Font
	// Color is the text color.
	Color color.RGBA
	// Hint contains the text displayed when the editor is empty.
	Hint string
	// HintColor is the color of hint text.
	HintColor color.RGBA
	// Color is the text color.
	TxColor color.RGBA
	BgColor color.RGBA
	shaper  text.Shaper
}

func (c DuoUIcounter) Layout(gtx *layout.Context) {
	th := material.NewTheme()

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
							//cs := gtx.Constraints

							in := layout.UniformInset(unit.Dp(0))
							in.Layout(gtx, func() {
								th.H3(fmt.Sprint(c.item.i)).Layout(gtx)
							})
						}),
						layout.Flexed(1, func() {
							layout.Flex{}.Layout(gtx,
								layout.Flexed(0.4, func() {
									c.increase.Layout(gtx)
								}),
								layout.Flexed(0.2, func() {
									c.reset.Layout(gtx)
								}),
								layout.Flexed(0.4, func() {
									c.decrease.Layout(gtx)
								}),
							)
						}),
					)

				}),
			)
		}),
	)
}

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

func counter(i *item) DuoUIcounter {
	return DuoUIcounter{
		item: i,
		increase: &button{
			Name: "increase",
			Do: func(n interface{}) {
				i.doIncrease(n.(int))
			},
			OperateValue: 1,
		},
		decrease: &button{
			Name: "decrease",
			Do: func(n interface{}) {
				i.doDecrease(n.(int))
			},
			OperateValue: 1,
		},
		reset: &button{
			Name: "reset",
			Do: func(interface{}) {
				i.doReset()
			},
			OperateValue: 0,
		},
	}
}

func main() {
	itemValue := &item{
		i: 0,
	}
	go func() {
		w := app.NewWindow()
		gofont.Register()
		c := counter(itemValue)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				c.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
