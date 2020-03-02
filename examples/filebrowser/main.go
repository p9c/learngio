package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"io/ioutil"
	"log"
)

func main() {
	//col := color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0x30}

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
		//listSub := &layout.List{
		//	Axis: layout.Vertical,
		//}
		//thingsCurrent := new(Thing)

		//selected := &Thing{
		//
		//	Name:     "/",
		//	pressed:  false,
		//	selected: false,
		//}

		//listRootThings := listThings("")
		//listSubThings, _ := readThings(selected)
		//things := make(map[string]*Thing)
		//for _, t := range listRootThings {
		//	things[t.Name] = new(Thing)
		//}
		//subThings := make(map[string]*Thing)
		//for _, t := range listSubThings {
		//	listRootThings[t.Name()] = new(Thing)
		//}

		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				widgets := []func(){
					func() {
						// Folder list
						//layout.List{
						//	Axis: layout.Vertical,
						//}.Layout(gtx, len(listRootThings), func(i int) {
						//	//button := new(Button)
						//	if listRootThings[i].IsDir() {
						//		col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
						//		things[listRootThings[i].Name].Layout(listRootThings[i].Name, col, gtx)
						//	} else {
						//		things[listRootThings[i].Name].Layout(listRootThings[i].Name, col, gtx)
						//	}
						//})
						///
					},
					func() {
						// Folder list
						listThings("/", gtx)
						//listSub.Layout(gtx, len(listSubThings), func(i int) {
						//	if listSubThings[i].IsDir() {
						//		col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
						//		subThings[listSubThings[i].Name()].Layout(listSubThings[i].Name(), col, gtx)
						//	} else {
						//		subThings[listSubThings[i].Name()].Layout(listSubThings[i].Name(), col, gtx)
						//	}
						//})
						//
					},
					func() {
						//selected.Info(selected, col, gtx)
					},
				}
				mainList.Layout(gtx, len(widgets), func(i int) {
					layout.UniformInset(unit.Dp(0)).Layout(gtx, widgets[i])
				})
				//////
				////

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

type Thing struct {
	Name     string
	Type     string
	out      interface{}
	pressed  bool
	selected bool
}

func (t *Thing) Info(col color.RGBA, gtx *layout.Context) {
	th := material.NewTheme()
	th.H6(t.Name).Layout(gtx)
}

// START OMIT
func (t *Thing) Layout(text string, col color.RGBA, gtx *layout.Context) {
	for _, e := range gtx.Events(t) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				t.pressed = true // HLevent
				t.selected = true
			case pointer.Release: // HLevent
				t.pressed = false // HLevent
			}
		}
	}
	th := material.NewTheme()

	if t.pressed {
		col = color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0xcf}
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: 500, Y: 500}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: t}.Add(gtx.Ops) // HLevent
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

func listThings(dir string, gtx *layout.Context) {
	list := &layout.List{
		Axis: layout.Vertical,
	}

	if dir == "" {
		dir = "/"
	}
	listThings, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	things := make(map[string]*Thing)
	for _, t := range listThings {
		things[t.Name()] = &Thing{
			Name: t.Name(),
		}
		if t.IsDir() {
			things[t.Name()].Type = "space"
		}
	}

	list.Layout(gtx, len(listThings), func(i int) {
		col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
		if listThings[i].IsDir() {
		} else {
			col = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
		}
		things[listThings[i].Name()].Layout(listThings[i].Name(), col, gtx)
	})
}
