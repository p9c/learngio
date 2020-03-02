package main

import (
	"fmt"

	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/font/gofont"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/widget/material"
)

func main() {
	gofont.Register()
	th := material.NewTheme()
	go func() {
		w := app.NewWindow()
		// START INIT OMIT
		list := &layout.List{
			Axis:        layout.Vertical,
			ScrollToEnd: true,
		}
		gtx := layout.NewContext(w.Queue())
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				drawList(gtx, list, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawList(gtx *layout.Context, list *layout.List, th *material.Theme) {
	const n = 1e6
	list.Layout(gtx, n, func(i int) {
		txt := fmt.Sprintf("List element #%d", i)

		th.H3(txt).Layout(gtx)
	})
}

// END OMIT
