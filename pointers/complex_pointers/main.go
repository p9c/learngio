package main

import (
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
	list := &layout.List{
		Axis: layout.Horizontal,
	}
	go func() {
		//NewWindow(options ...Option) *Window {
		w := app.NewWindow(
			app.Size(unit.Dp(900), unit.Dp(556)),
			app.Title("Learn Gio"),
		)

		gofont.Register()
		//th := material.NewTheme()

		file := &Button{
			Name: "File",
			Do: func(interface{}) {
				itemValue.doReset()
			},
			OperateValue:   0,
			ColorBg:        "ff303030",
			ColorBgHover:   "ffcf3030",
			ColorText:      "ffcfcfcf",
			ColorTextHover: "ffcfcfcf",
		}
		edit := &Button{
			Name: "Edit",
			Do: func(interface{}) {
				itemValue.doReset()
			},
			OperateValue:   0,
			ColorBg:        "ff303030",
			ColorBgHover:   "ffcf3030",
			ColorText:      "ffcfcfcf",
			ColorTextHover: "ffcfcfcf",
		}

		help := &Button{
			Name: "Help",
			Do: func(interface{}) {
				itemValue.doReset()
			},
			OperateValue:   0,
			ColorBg:        "ff303030",
			ColorBgHover:   "ffcf3030",
			ColorText:      "ffcfcfcf",
			ColorTextHover: "ffcfcfcf",
		}

		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

						in := layout.UniformInset(unit.Dp(0))
						in.Layout(gtx, func() {
							layout.Align(layout.Start).Layout(gtx, func() {
								//th.H6(fmt.Sprint(itemValue.i)).Layout(gtx)
								widgets := []func(){
									func() {
										layout.Flex{}.Layout(gtx,
											layout.Rigid(func() {
												file.Layout(gtx)
											}),
											layout.Rigid(func() {
												edit.Layout(gtx)
											}),
											layout.Rigid(func() {
												help.Layout(gtx)
											}),
										)

									},

									//func() {
									//	th.H6(fmt.Sprint(itemValue.i)).Layout(gtx)
									//},
								}
								list.Layout(gtx, len(widgets), func(i int) {
									layout.UniformInset(unit.Dp(0)).Layout(gtx, widgets[i])
								})

							})
						})
					}),
					layout.Flexed(1, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ffcfcfcf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

type Button struct {
	pressed        bool
	Name           string
	Do             func(interface{})
	ColorBg        string
	ColorBgHover   string
	ColorText      string
	ColorTextHover string
	BorderRadius   [4]float32
	OperateValue   interface{}
}

func (b *Button) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
				//b.Do(b.OperateValue)
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}

	cs := gtx.Constraints
	th := material.NewTheme()

	//colorBg := helpers.HexARGB("ff30cfcf")
	colorText := helpers.HexARGB(b.ColorText)
	colorBg := helpers.HexARGB(b.ColorBg)

	if b.pressed {
		colorText = helpers.HexARGB(b.ColorTextHover)
		colorBg = helpers.HexARGB(b.ColorBgHover)
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))

	in := layout.UniformInset(unit.Dp(5))
	in.Layout(gtx, func() {
		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))
		//cs := gtx.Constraints
		layout.Align(layout.Center).Layout(gtx, func() {
			txt := th.Caption(b.Name)
			txt.Color = colorText
			txt.Layout(gtx)
		})
	})
}
