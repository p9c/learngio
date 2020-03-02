package main

import (
	"fmt"
	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/f32"
	"github.com/gop9/olt/gio/font/gofont"
	"github.com/gop9/olt/gio/io/pointer"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/op/paint"
	"github.com/gop9/olt/gio/widget/material"
	"image"
	"image/color"
)

var text = "button txt"

// START QUEUE OMIT
func main() {
	go func() {
		w := app.NewWindow()
		gofont.Register()

		surface := new(Surface)
		//th := material.NewTheme()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				surface.Layout(gtx)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// END QUEUE OMIT

type Surface struct {
	pressed bool
}

// START OMIT
func (b *Surface) Layout(gtx *layout.Context) {
	th := material.NewTheme()
	cs := gtx.Constraints
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			text = fmt.Sprint(e.Position)
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
				text = fmt.Sprint(e.Position)
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
				text = fmt.Sprint(e.Position)
			}
		}
	}

	col := color.RGBA{A: 0xff, R: 0xff}
	if b.pressed {
		col = color.RGBA{A: 0xff, G: 0xff}
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	layout.Flex{}.Layout(gtx,
		layout.Rigid(func() {
			square := f32.Rectangle{
				Max: f32.Point{X: float32(cs.Width.Max), Y: float32(cs.Height.Max)},
			}
			paint.ColorOp{Color: col}.Add(gtx.Ops)
			paint.PaintOp{Rect: square}.Add(gtx.Ops)
		}),
		layout.Rigid(func() {
			th.H6(text).Layout(gtx)
		}))

}

// END OMIT
