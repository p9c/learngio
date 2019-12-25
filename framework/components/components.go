package main

import (
	"flag"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/p9c/learngio/framework"
	"github.com/p9c/learngio/helpers"
	"log"
)

var (
	list = &layout.List{
		Axis: layout.Vertical,
	}
	button = new(widget.Button)
)

func main() {
	flag.Parse()
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
	fwk := framework.NewFramework()
	gtx := layout.NewContext(w.Queue())

	//buttonComponent := &framework.Button{
	//	Name: "increase",
	//	Do: func(interface{}) {
	//	},
	//	BorderRadius: [4]float32{30, 30, 30, 30},
	//}
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			cs := gtx.Constraints
			helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("30303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

			widgets := []func(){
				func() {
					//th.H3(topLabel).Layout(gtx)
				},
				func() {
					gtx.Constraints.Height.Max = gtx.Px(unit.Dp(120))
					fwk.Button("sss").Layout(gtx, button)
				},
			}
			list.Layout(gtx, len(widgets), func(i int) {
				layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
			})
			e.Frame(gtx.Ops)
		}
	}
}
