package main

import (
	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/f32"
	"github.com/gop9/olt/gio/font/gofont"
	"github.com/gop9/olt/gio/io/pointer"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/op"
	"github.com/gop9/olt/gio/op/paint"
	"github.com/gop9/olt/gio/unit"
	"github.com/gop9/olt/gio/widget/material"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"
)

var (
	seleceted string
)

func main() {
	listRootThings, _ := readThings("")
	listSubThings, _ := readThings(seleceted)
	col := color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0x30}
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gofont.Register()
		// START INIT OMIT
		//tabs := &[]Tab{
		//	Tab{
		//		name: "thing",
		//	},
		//	Tab{
		//		name: "info",
		//	},
		//}
		mainList := &layout.List{
			Axis: layout.Horizontal,
		}
		//listRoot := &layout.List{
		//	Axis: layout.Vertical,
		//}
		listSub := &layout.List{
			Axis: layout.Vertical,
		}

		things := *new(map[string]*Button)

		for _, t := range listRootThings {
			things[t.Name()] = new(Button)
		}
		//button := new(Button)

		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				widgets := []func(){
					//func() {
					//	// Folder list
					//	listRoot.Layout(gtx, len(listRootThings), func(i int) {
					//		//button := new(Button)
					//		if listRootThings[i].IsDir() {
					//			col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
					//			button.Layout(listRootThings[i].Name(), col, gtx)
					//		} else {
					//			button.Layout(listRootThings[i].Name(), col, gtx)
					//		}
					//	})
					//	///
					//},
					func() {
						// Folder list
						listSub.Layout(gtx, len(listSubThings), func(i int) {
							if listSubThings[i].IsDir() {
								col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
								things[listSubThings[i].Name()].Layout(listSubThings[i].Name(), col, gtx)
							} else {
								things[listSubThings[i].Name()].Layout(listSubThings[i].Name(), col, gtx)
							}
						})
						///
					},
				}
				mainList.Layout(gtx, len(widgets), func(i int) {
					layout.UniformInset(unit.Dp(3)).Layout(gtx, widgets[i])
				})
				//////
				////

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

type Tab struct {
	name   string
	layout layout.List
}

type Button struct {
	pressed bool
}

// START OMIT
func (b *Button) Layout(text string, col color.RGBA, gtx *layout.Context) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
				seleceted = text
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}
	th := material.NewTheme()

	if b.pressed {
		col = color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0xcf}
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: 500, Y: 500}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	drawSquare(gtx.Ops, col)
	th.H6(text).Layout(gtx)
}

// END OMIT

func drawSquare(ops *op.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
}

func readThings(dir string) (files []os.FileInfo, err error) {
	if dir == "" {
		dir = "/"
	}
	files, err = ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return
}
