// SPDX-License-Identifier: Unlicense OR MIT

package main

// A Gio program that demonstrates Gio widgets. See https://gioui.org for more information.

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
	"image/color"
	"log"

	"gioui.org/font/gofont"
)

type App struct {
	Position     float32
	Cursor       f32.Point
	OperateValue interface{}
	Height       int
	CursorHeight float32
	Icon         material.Icon
}

func main() {

	a := &App{}

	ic, err := material.NewIcon(icons.ContentAdd)
	if err != nil {
		log.Fatal(err)
	}
	icon = ic
	gofont.Register()

	go func() {
		w := app.NewWindow()
		if err := a.loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func (a *App) loop(w *app.Window) error {
	//th := material.NewTheme()
	gtx := layout.NewContext(w.Queue())
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			for _, ew := range gtx.Events(a) {
				if ew, ok := ew.(pointer.Event); ok {
					//a.Position = ew.Position.Y - (a.CursorHeight / 2)
					switch ew.Type {
					case pointer.Press:
						//p.scrollBar.body.pressed = true
						//p.scrollBar.body.Do(p.scrollBar.body.OperateValue)
						//list.Position.First = int(s.Position)
					case pointer.Release:
						//p.scrollBar.body.pressed = false
					}
					//fmt.Println(ew.Position)
					//fmt.Println(ew.Scroll)
					a.Cursor = ew.Position
				}
			}
			cs := gtx.Constraints
			pointer.Rect(
				image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}},
			).Add(gtx.Ops)
			pointer.InputOp{Key: a}.Add(gtx.Ops)

			uni := layout.UniformInset(unit.Dp(0))
			uni.Top = unit.Dp(a.Cursor.Y)
			uni.Left = unit.Dp(a.Cursor.X)
			uni.Layout(gtx, func() {
				square := f32.Rectangle{
					Max: f32.Point{
						X: float32(33),
						Y: float32(33),
					},
				}
				paint.ColorOp{Color: color.RGBA{0x00, 0x00, 0xff, 0xff}}.Add(gtx.Ops)
				clip.Rect{Rect: square}.Op(gtx.Ops).Add(gtx.Ops)
				paint.PaintOp{Rect: square}.Add(gtx.Ops)
				gtx.Dimensions = layout.Dimensions{Size: image.Point{X: 33, Y: 33}}
			})
			e.Frame(gtx.Ops)
		}
	}
}

var (
	icon *material.Icon
)
