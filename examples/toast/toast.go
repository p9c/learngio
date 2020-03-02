package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

var (
	text       = "button txt"
	buttonOpen = new(widget.Button)
)

// START QUEUE OMIT
func main() {
	go func() {
		w := app.NewWindow()
		//toast := new(Toast)
		//th := material.NewTheme()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				//layout.Flex{}.Layout(gtx,
				//	layout.Rigid(func() {
				var button material.Button
				button.Layout(gtx, buttonOpen)
				//}),
				//layout.Rigid(func() {
				//toast.Layout(gtx)
				//}),
				//)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// END QUEUE OMIT

//type Button struct {
//	pressed bool
//}
//
//func (b *Button) Layout(gtx *layout.Context) {
//	for _, e := range gtx.Events(b) { // HLevent
//		if e, ok := e.(pointer.Event); ok { // HLevent
//			switch e.Type { // HLevent
//			case pointer.Press: // HLevent
//				b.pressed = true // HLevent
//			case pointer.Release: // HLevent
//				b.pressed = false // HLevent
//			}
//		}
//	}
//
//	col := color.RGBA{A: 0xff, R: 0xff}
//	if b.pressed {
//		col = color.RGBA{A: 0xff, G: 0xff}
//	}
//	pointer.Rect( // HLevent
//		image.Rectangle{Max: image.Point{X: 50, Y: 50}}, // HLevent
//	).Add(gtx.Ops)                       // HLevent
//	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
//	square := f32.Rectangle{
//		Max: f32.Point{X: 50, Y: 50},
//	}
//	paint.ColorOp{Color: col}.Add(gtx.Ops)
//	paint.PaintOp{Rect: square}.Add(gtx.Ops)
//	//th.H6(text).Layout(gtx)
//
//}

type Toast struct {
	hover   bool
	pressed bool
}

func (t *Toast) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(t) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				t.pressed = true // HLevent
			case pointer.Release: // HLevent
				t.pressed = false // HLevent
			}

		}
	}

	col := color.RGBA{A: 0xff, B: 0xff}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: 50, Y: 50}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: t}.Add(gtx.Ops) // HLevent
	square := f32.Rectangle{
		Max: f32.Point{X: 50, Y: 50},
	}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: square}.Add(gtx.Ops)
	//th.H6(text).Layout(gtx)
}
