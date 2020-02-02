package main

import (
	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/font/gofont"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/unit"
	"github.com/gop9/olt/gio/widget/material"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("/")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gofont.Register()
		th := material.NewTheme()
		// START INIT OMIT
		list := &layout.List{
			Axis: layout.Vertical,
		}
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				list.Layout(gtx, len(files), func(i int) {
					if files[i].IsDir() {
						th.H4("/" + files[i].Name()).Layout(gtx)
					} else {
						th.H6("/" + files[i].Name()).Layout(gtx)
					}
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// END OMIT
//
//layout.Flex{
//	Axis: layout.Horizontal,
//}.Layout(gtx,
//	layout.Flexed(0.5, func() {
//		cs := gtx.Constraints
//		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ffcf30cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
//	}),
//	layout.Flexed(0.25, func() {
//		cs := gtx.Constraints
//		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
//	}),
//	layout.Flexed(0.25, func() {
//		cs := gtx.Constraints
//		helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
//	}),
//)
