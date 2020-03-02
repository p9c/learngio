package main

import (
	"flag"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/examples/html/native/controller"
	"github.com/p9c/learngio/examples/html/native/model"
	"github.com/p9c/learngio/helpers"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
	"image/color"
	"log"
)

var url = `https://parallelcoin.io`

var (
	editor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	list = &layout.List{
		Axis: layout.Vertical,
	}
	icon *material.Icon
)

func main() {

	flag.Parse()
	gofont.Register()
	th := material.NewTheme()

	editor.SetText(url)
	ic, err := material.NewIcon(icons.ContentAdd)
	if err != nil {
		log.Fatal(err)
	}
	icon = ic

	resource := controller.GetResource(url)
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(1400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(header(gtx, th)),
					layout.Flexed(1, body(gtx, th, resource)),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func drawRectangle(gtx *layout.Context, w, h int, color color.RGBA, borderRadius [4]float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		//cs := gtx.Constraints
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(w),
				Y: float32(h),
			},
		}
		paint.ColorOp{Color: color}.Add(gtx.Ops)

		clip.Rect{Rect: square,
			NE: borderRadius[0], NW: borderRadius[1], SE: borderRadius[2], SW: borderRadius[3]}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: w, Y: h}}
	})
}

func header(gtx *layout.Context, th *material.Theme) func() {
	return func() {
		cs := gtx.Constraints
		drawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ffcf30cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
		layout.Center.Layout(gtx,
			func() {
				dr := func() {
					//gtx.Constraints.Height.Max = gtx.Px(unit.Dp(200))
					e := th.Editor(url)
					e.Layout(gtx, editor)
					for _, e := range editor.Events(gtx) {
						if e, ok := e.(widget.SubmitEvent); ok {
							url = e.Text
							editor.SetText("")
						}
					}
				}
				layout.UniformInset(unit.Dp(16)).Layout(gtx, dr)
			})
	}
}

func body(gtx *layout.Context, th *material.Theme, resource *model.Resource) func() {
	return func() {
		cs := gtx.Constraints
		drawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

		parsed := controller.ParseDocument(resource.Body)
		controller.RenderDocument(gtx, th, parsed)

	}
}
