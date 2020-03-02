// SPDX-License-Identifier: Unlicense OR MIT

package main

// A Gio program that demonstrates Gio widgets. See https://gioui.org for more information.

import (
	"flag"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"log"

	"gioui.org/font/gofont"
)

func main() {
	flag.Parse()
	editor.SetText(longText)
	ic, err := material.NewIcon(icons.ContentAdd)
	if err != nil {
		log.Fatal(err)
	}
	icon = ic
	gofont.Register()

	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()
	gtx := layout.NewContext(w.Queue())
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)

			dr := func() {
				//gtx.Constraints.Height.Max = gtx.Px(unit.Dp(200))
				th.Editor("").Layout(gtx, editor)
			}
			layout.UniformInset(unit.Dp(16)).Layout(gtx, dr)
			e.Frame(gtx.Ops)
		}
	}
}

var (
	editor = new(widget.Editor)
	list   = &layout.List{
		Axis: layout.Vertical,
	}
	icon *material.Icon
)

const longText = `package main

func main(){
	fmt.Println("Functions inside funcion of huge fuction which is method")
}
`
